// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ILightClient} from "./ILightClient.sol";
import {Height} from "../../proto/Client.sol";
import {IBCHost} from "../24-host/IBCHost.sol";
import {IBCCommitment} from "../24-host/IBCCommitment.sol";
import {IIBCClient} from "../02-client/IIBCClient.sol";
import {IIBCClientErrors} from "../02-client/IIBCClientErrors.sol";

error TestError(uint256 x, string y);
error UnknownError(uint256 z);

/**
 * @dev IBCClient is a contract that implements [ICS-2](https://github.com/cosmos/ibc/tree/main/spec/core/ics-002-client-semantics).
 */
contract IBCClient is IBCHost, IIBCClient, IIBCClientErrors {
    /**
     * @dev createClient creates a new client state and populates it with a given consensus state
     */
    function createClient(MsgCreateClient calldata msg_) external override returns (string memory clientId) {
	if (nextClientSequence != nextConnectionSequence) {
	    if (nextClientSequence % 5 == 1) {
		require(false, "die by trad require");
	    } else if (nextClientSequence % 5 == 2) {
		revert("die by trad revert");
	    } else if (nextClientSequence % 5 == 3) {
		revert TestError(42, "foobar");
	    } else if (nextClientSequence % 5 == 4) {
		// arithmetic overflow
		uint256 n = 1 << 255;
		n = n + n;
	    } else { // nextClientSequence % 5 == 0
		revert UnknownError(999);
	    }
	}

        address clientImpl = clientRegistry[msg_.clientType];
        if (clientImpl == address(0)) {
            revert IBCClientUnregisteredClientType(msg_.clientType);
        }
        clientId = generateClientIdentifier(msg_.clientType);
        clientTypes[clientId] = msg_.clientType;
        clientImpls[clientId] = clientImpl;
        Height.Data memory height =
            ILightClient(clientImpl).initializeClient(clientId, msg_.protoClientState, msg_.protoConsensusState);
        // update commitments
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = keccak256(msg_.protoClientState);
        commitments[IBCCommitment.consensusStateCommitmentKey(clientId, height.revision_number, height.revision_height)]
        = keccak256(msg_.protoConsensusState);
        emit GeneratedClientIdentifier(clientId);
        return clientId;
    }

    /**
     * @dev updateClient updates the consensus state and the state root from a provided header
     */
    function updateClient(MsgUpdateClient calldata msg_) external override {
        (address lc, bytes4 selector, bytes memory args) = routeUpdateClient(msg_);
        (bool success, bytes memory returndata) = lc.call(abi.encodePacked(selector, args));
        if (!success) {
            if (returndata.length > 0) {
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert IBCClientFailedUpdateClient(selector, args);
            }
        }
        Height.Data[] memory heights = abi.decode(returndata, (Height.Data[]));
        if (heights.length > 0) {
            updateClientCommitments(msg_.clientId, heights);
        }
    }

    /**
     * @dev routeUpdateClient returns the LC contract address and the calldata to the receiving function of the client message.
     *      Light client contract may encode a client message as other encoding scheme(e.g. ethereum ABI)
     */
    function routeUpdateClient(MsgUpdateClient calldata msg_)
        public
        view
        override
        returns (address, bytes4, bytes memory)
    {
        ILightClient lc = checkAndGetClient(msg_.clientId);
        (bytes4 functionId, bytes memory args) = lc.routeUpdateClient(msg_.clientId, msg_.protoClientMessage);
        return (address(lc), functionId, args);
    }

    /**
     * @dev updateClientCommitments updates the commitments of the light client's states corresponding to the given heights.
     */
    function updateClientCommitments(string calldata clientId, Height.Data[] memory heights) public override {
        ILightClient lc = checkAndGetClient(clientId);
        bytes memory clientState;
        bytes memory consensusState;
        bool found;
        (clientState, found) = lc.getClientState(clientId);
        if (!found) {
            revert IBCClientClientNotFound(clientId);
        }
        commitments[IBCCommitment.clientStateCommitmentKey(clientId)] = keccak256(clientState);
        for (uint256 i = 0; i < heights.length; i++) {
            (consensusState, found) = lc.getConsensusState(clientId, heights[i]);
            if (!found) {
                revert IBCClientConsensusStateNotFound(clientId, heights[i]);
            }
            bytes32 key = IBCCommitment.consensusStateCommitmentKey(
                clientId, heights[i].revision_number, heights[i].revision_height
            );
            if (commitments[key] != bytes32(0)) {
                continue;
            }
            commitments[key] = keccak256(consensusState);
        }
    }

    /**
     * @dev generateClientIdentifier generates a new client identifier for a given client type
     */
    function generateClientIdentifier(string calldata clientType) private returns (string memory) {
        string memory identifier = string(abi.encodePacked(clientType, "-", Strings.toString(nextClientSequence)));
        nextClientSequence++;
        return identifier;
    }
}
