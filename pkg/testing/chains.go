package testing

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibcidentifier"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20bank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20transferbank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/simpletoken"
	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/channel"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	ibft2clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client/ibft2"
	mockclienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client/mock"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/irohaeth"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/wallet"
)

const (
	DefaultChannelVersion        = "ics20-1"
	BlockTime             uint64 = 1000 * 1000 * 1000 // 1[sec]
	DefaultDelayPeriod    uint64 = 3 * BlockTime
	DefaultPrefix                = "ibc"
	TransferPort                 = "transfer"

	RelayerKeyIndex uint32 = 0
)

var (
	abiHandlerContract abi.ABI
	abiSendPacket,
	abiGeneratedClientIdentifier,
	abiGeneratedConnectionIdentifier,
	abiGeneratedChannelIdentifier abi.Event
)

func init() {
	var err error
	abiHandlerContract, err = abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		panic(err)
	}
	parsedHostABI, err := abi.JSON(strings.NewReader(ibchost.IbchostABI))
	if err != nil {
		panic(err)
	}
	abiSendPacket = abiHandlerContract.Events["SendPacket"]
	abiGeneratedClientIdentifier = parsedHostABI.Events["GeneratedClientIdentifier"]
	abiGeneratedConnectionIdentifier = parsedHostABI.Events["GeneratedConnectionIdentifier"]
	abiGeneratedChannelIdentifier = parsedHostABI.Events["GeneratedChannelIdentifier"]
}

type Chain struct {
	t *testing.T

	// Core Modules
	client        client.Client
	IBCHandler    ibchandler.Ibchandler
	IBCHost       ibchost.Ibchost
	IBCIdentifier ibcidentifier.Ibcidentifier

	// App Modules
	SimpleToken   simpletoken.Simpletoken
	ICS20Transfer ics20transferbank.Ics20transferbank
	ICS20Bank     ics20bank.Ics20bank

	chainID int64

	ContractConfig ContractConfig

	mnemonicPhrase string
	keys           map[uint32]*ecdsa.PrivateKey

	// State
	LastContractState client.ContractState

	// IBC specific helpers
	ClientIDs   []string          // ClientID's used on this chain
	Connections []*TestConnection // track connectionID's created for this chain
	IBCID       uint64

	// Iroha specific items
	irohaEthClient     *irohaeth.Client
	ibcHandlerContract *irohaeth.BoundContract
}

type ContractConfig interface {
	GetIBCHostAddress() common.Address
	GetIBCHandlerAddress() common.Address
	GetIBCIdentifierAddress() common.Address
	GetIBFT2ClientAddress() common.Address
	GetMockClientAddress() common.Address

	GetSimpleTokenAddress() common.Address
	GetICS20TransferBankAddress() common.Address
	GetICS20BankAddress() common.Address
}

func NewChain(t *testing.T, chainID int64, client client.Client, config ContractConfig, mnemonicPhrase string, ibcID uint64) *Chain {
	ibcHost, err := ibchost.NewIbchost(config.GetIBCHostAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ibcHandler, err := ibchandler.NewIbchandler(config.GetIBCHandlerAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ibcIdentifier, err := ibcidentifier.NewIbcidentifier(config.GetIBCIdentifierAddress(), client)
	if err != nil {
		t.Error(err)
	}
	simpletoken, err := simpletoken.NewSimpletoken(config.GetSimpleTokenAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ics20transfer, err := ics20transferbank.NewIcs20transferbank(config.GetICS20TransferBankAddress(), client)
	if err != nil {
		t.Error(err)
	}
	ics20bank, err := ics20bank.NewIcs20bank(config.GetICS20BankAddress(), client)
	if err != nil {
		t.Error(err)
	}

	irohaEthClient := irohaeth.NewClient(client.Conn)
	ibcHandlerContract := irohaeth.NewBoundContract(config.GetIBCHandlerAddress(), abiHandlerContract, irohaEthClient)

	return &Chain{
		t:              t,
		client:         client,
		chainID:        chainID,
		ContractConfig: config,
		mnemonicPhrase: mnemonicPhrase,
		keys:           make(map[uint32]*ecdsa.PrivateKey),
		IBCID:          ibcID,

		IBCHost:       *ibcHost,
		IBCHandler:    *ibcHandler,
		IBCIdentifier: *ibcIdentifier,
		SimpleToken:   *simpletoken,
		ICS20Transfer: *ics20transfer,
		ICS20Bank:     *ics20bank,

		irohaEthClient:     irohaEthClient,
		ibcHandlerContract: ibcHandlerContract,
	}
}

func (chain *Chain) Client() client.Client {
	return chain.client
}

func (chain *Chain) ClientType() string {
	return chain.client.ClientType()
}

func (chain *Chain) TxOpts(ctx context.Context, index uint32) *bind.TransactOpts {
	return client.MakeGenTxOpts(big.NewInt(chain.chainID), chain.prvKey(index))(ctx)
}

func (chain *Chain) CallOpts(ctx context.Context, index uint32) *bind.CallOpts {
	opts := chain.TxOpts(ctx, index)
	return &bind.CallOpts{
		From:    opts.From,
		Context: opts.Context,
	}
}

func (chain *Chain) transactAndWaitMined(opts *bind.TransactOpts, method string, params ...interface{}) error {
	tx, err := chain.ibcHandlerContract.Transact(opts, method, params...)
	if err != nil {
		return fmt.Errorf("failed to transact: %v", err)
	}
	rc, err := chain.irohaEthClient.WaitMined(opts.Context, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for tx to be mined: %v", err)
	} else if rc.Status != 1 {
		return fmt.Errorf("transaction status error: %v", rc)
	}
	return nil
}

func (chain *Chain) prvKey(index uint32) *ecdsa.PrivateKey {
	key, ok := chain.keys[index]
	if ok {
		return key
	}
	key, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(chain.mnemonicPhrase, fmt.Sprintf("m/44'/60'/0'/0/%v", index))
	if err != nil {
		panic(err)
	}
	chain.keys[index] = key
	return key
}

func (chain *Chain) ChainID() int64 {
	return chain.chainID
}

func (chain *Chain) ChainIDString() string {
	return fmt.Sprint(chain.chainID)
}

func (chain *Chain) GetCommitmentPrefix() []byte {
	return []byte(DefaultPrefix)
}

func (chain *Chain) GetIBFT2ClientState(clientID string) *ibft2clienttypes.ClientState {
	ctx := context.Background()
	bz, found, err := chain.IBCHost.GetClientState(chain.CallOpts(ctx, RelayerKeyIndex), clientID)
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("clientState not found")
	}
	var cs ibft2clienttypes.ClientState
	if err := UnmarshalWithAny(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetMockClientState(clientID string) *mockclienttypes.ClientState {
	ctx := context.Background()
	bz, found, err := chain.IBCHost.GetClientState(chain.CallOpts(ctx, RelayerKeyIndex), clientID)
	if err != nil {
		require.NoError(chain.t, err)
	} else if !found {
		panic("clientState not found")
	}
	var cs mockclienttypes.ClientState
	if err := UnmarshalWithAny(bz, &cs); err != nil {
		panic(err)
	}
	return &cs
}

func (chain *Chain) GetContractState(counterparty *Chain, counterpartyClientID string, storageKeys [][]byte, height *big.Int) (client.ContractState, error) {
	if height == nil {
		switch counterparty.ClientType() {
		case ibcclient.MockClient:
			height = counterparty.GetMockClientState(counterpartyClientID).LatestHeight.ToBN()
		case ibcclient.BesuIBFT2Client:
			height = counterparty.GetIBFT2ClientState(counterpartyClientID).LatestHeight.ToBN()
		default:
			return nil, fmt.Errorf("unknown client type: '%v'", counterparty.ClientType())
		}
	}
	return chain.client.GetContractState(
		context.Background(),
		chain.ContractConfig.GetIBCHostAddress(),
		storageKeys,
		height,
	)
}

func (chain *Chain) ConstructMockMsgCreateClient(counterparty *Chain) ibchandler.IBCMsgsMsgCreateClient {
	clientState := mockclienttypes.ClientState{
		LatestHeight: ibcclient.NewHeightFromBN(counterparty.LastHeader().Number),
	}
	consensusState := mockclienttypes.ConsensusState{
		Timestamp: counterparty.LastHeader().Time,
	}
	clientStateBytes, err := MarshalWithAny(&clientState)
	if err != nil {
		panic(err)
	}
	consensusStateBytes, err := MarshalWithAny(&consensusState)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          ibcclient.MockClient,
		Height:              clientState.LatestHeight.ToCallData(),
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	}
}

func (chain *Chain) ConstructIBFT2MsgCreateClient(counterparty *Chain) ibchandler.IBCMsgsMsgCreateClient {
	clientState := ibft2clienttypes.ClientState{
		ChainId:         counterparty.ChainIDString(),
		IbcStoreAddress: counterparty.ContractConfig.GetIBCHostAddress().Bytes(),
		LatestHeight:    ibcclient.NewHeightFromBN(counterparty.LastHeader().Number),
	}
	consensusState := ibft2clienttypes.ConsensusState{
		Timestamp:  counterparty.LastHeader().Time,
		Root:       counterparty.LastHeader().Root.Bytes(),
		Validators: counterparty.LastContractState.(client.IBFT2ContractState).Validators(),
	}
	clientStateBytes, err := MarshalWithAny(&clientState)
	if err != nil {
		panic(err)
	}
	consensusStateBytes, err := MarshalWithAny(&consensusState)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          ibcclient.BesuIBFT2Client,
		Height:              clientState.LatestHeight.ToCallData(),
		ClientStateBytes:    clientStateBytes,
		ConsensusStateBytes: consensusStateBytes,
	}
}

func (chain *Chain) ConstructMockMsgUpdateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgUpdateClient {
	cs := counterparty.LastContractState.(client.ETHContractState)
	header := mockclienttypes.Header{
		Height:    ibcclient.NewHeightFromBN(cs.Header().Number),
		Timestamp: cs.Header().Time,
	}
	bz, err := MarshalWithAny(&header)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: clientID,
		Header:   bz,
	}
}

func (chain *Chain) ConstructIBFT2MsgUpdateClient(counterparty *Chain, clientID string) ibchandler.IBCMsgsMsgUpdateClient {
	trustedHeight := chain.GetIBFT2ClientState(clientID).LatestHeight
	cs := counterparty.LastContractState.(client.IBFT2ContractState)
	var header = ibft2clienttypes.Header{
		BesuHeaderRlp:     cs.SealingHeaderRLP(),
		Seals:             cs.CommitSeals,
		TrustedHeight:     trustedHeight,
		AccountStateProof: cs.ETHProof().AccountProofRLP,
	}
	bz, err := MarshalWithAny(&header)
	if err != nil {
		panic(err)
	}
	return ibchandler.IBCMsgsMsgUpdateClient{
		ClientId: clientID,
		Header:   bz,
	}
}

func (chain *Chain) UpdateHeader() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		state, err := chain.client.GetContractState(ctx, chain.ContractConfig.GetIBCHostAddress(), nil, nil)
		if err != nil {
			panic(err)
		}
		if chain.LastContractState == nil || state.Header().Number.Cmp(chain.LastHeader().Number) == 1 {
			chain.LastContractState = state
			return
		} else {
			continue
		}
	}
}

func (chain *Chain) CreateMockClient(ctx context.Context, counterparty *Chain) (string, error) {
	msg := chain.ConstructMockMsgCreateClient(counterparty)
	if err := chain.transactAndWaitMined(chain.TxOpts(ctx, RelayerKeyIndex), "createClient", msg); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedClientID(ctx)
}

func (chain *Chain) UpdateMockClient(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructMockMsgUpdateClient(counterparty, clientID)
	return chain.transactAndWaitMined(chain.TxOpts(ctx, RelayerKeyIndex), "updateClient", msg)
}

func (chain *Chain) CreateIBFT2Client(ctx context.Context, counterparty *Chain) (string, error) {
	msg := chain.ConstructIBFT2MsgCreateClient(counterparty)
	if err := chain.WaitIfNoError(ctx)(
		chain.IBCHandler.CreateClient(chain.TxOpts(ctx, RelayerKeyIndex), msg),
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedClientID(ctx)
}

func (chain *Chain) UpdateIBFT2Client(ctx context.Context, counterparty *Chain, clientID string) error {
	msg := chain.ConstructIBFT2MsgUpdateClient(counterparty, clientID)
	return chain.WaitIfNoError(ctx)(
		chain.IBCHandler.UpdateClient(chain.TxOpts(ctx, RelayerKeyIndex), msg),
	)
}

func (chain *Chain) ConnectionOpenInit(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) (string, error) {
	if err := chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"connectionOpenInit",
		ibchandler.IBCMsgsMsgConnectionOpenInit{
			ClientId: connection.ClientID,
			Counterparty: ibchandler.CounterpartyData{
				ClientId:     connection.CounterpartyClientID,
				ConnectionId: "",
				Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
			},
			DelayPeriod: DefaultDelayPeriod,
		},
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedConnectionID(ctx)
}

func (chain *Chain) ConnectionOpenTry(ctx context.Context, counterparty *Chain, connection, counterpartyConnection *TestConnection) (string, error) {
	proofConnection, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return "", err
	}
	clientStateBytes, proofClient, err := counterparty.QueryClientProof(chain, counterpartyConnection.ClientID, proofConnection.Height.ToBN())
	if err != nil {
		return "", err
	}
	if err := chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"connectionOpenTry",
		ibchandler.IBCMsgsMsgConnectionOpenTry{
			PreviousConnectionId: "",
			Counterparty: ibchandler.CounterpartyData{
				ClientId:     counterpartyConnection.ClientID,
				ConnectionId: counterpartyConnection.ID,
				Prefix:       ibchandler.MerklePrefixData{KeyPrefix: counterparty.GetCommitmentPrefix()},
			},
			DelayPeriod:      DefaultDelayPeriod,
			ClientId:         connection.ClientID,
			ClientStateBytes: clientStateBytes,
			CounterpartyVersions: []ibchandler.VersionData{
				{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
			},
			ProofHeight: proofConnection.Height.ToCallData(),
			ProofInit:   proofConnection.Data,
			ProofClient: proofClient.Data,
		},
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedConnectionID(ctx)
}

// ConnectionOpenAck will construct and execute a MsgConnectionOpenAck.
func (chain *Chain) ConnectionOpenAck(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proofConnection, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return err
	}
	clientStateBytes, proofClient, err := counterparty.QueryClientProof(chain, counterpartyConnection.ClientID, proofConnection.Height.ToBN())
	if err != nil {
		return err
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"connectionOpenAck",
		ibchandler.IBCMsgsMsgConnectionOpenAck{
			ConnectionId:             connection.ID,
			CounterpartyConnectionID: counterpartyConnection.ID,
			ClientStateBytes:         clientStateBytes,
			Version:                  ibchandler.VersionData{Identifier: "1", Features: []string{"ORDER_ORDERED", "ORDER_UNORDERED"}},
			ProofHeight:              proofConnection.Height.ToCallData(),
			ProofTry:                 proofConnection.Data,
			ProofClient:              proofClient.Data,
		},
	)
}

func (chain *Chain) ConnectionOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	connection, counterpartyConnection *TestConnection,
) error {
	proof, err := counterparty.QueryConnectionProof(chain, connection.ClientID, counterpartyConnection.ID, nil)
	if err != nil {
		return err
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"connectionOpenConfirm",
		ibchandler.IBCMsgsMsgConnectionOpenConfirm{
			ConnectionId: connection.ID,
			ProofAck:     proof.Data,
			ProofHeight:  proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) ChannelOpenInit(
	ctx context.Context,
	ch, counterparty TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) (string, error) {
	if err := chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelOpenInit",
		ibchandler.IBCMsgsMsgChannelOpenInit{
			PortId: ch.PortID,
			Channel: ibchandler.ChannelData{
				State:    uint8(channeltypes.INIT),
				Ordering: uint8(order),
				Counterparty: ibchandler.ChannelCounterpartyData{
					PortId:    counterparty.PortID,
					ChannelId: "",
				},
				ConnectionHops: []string{connectionID},
				Version:        ch.Version,
			},
		},
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedChannelID(ctx)
}

func (chain *Chain) ChannelOpenTry(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	order channeltypes.Channel_Order,
	connectionID string,
) (string, error) {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return "", err
	}
	if err := chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelOpenTry",
		ibchandler.IBCMsgsMsgChannelOpenTry{
			PortId: ch.PortID,
			Channel: ibchandler.ChannelData{
				State:    uint8(channeltypes.TRYOPEN),
				Ordering: uint8(order),
				Counterparty: ibchandler.ChannelCounterpartyData{
					PortId:    counterpartyCh.PortID,
					ChannelId: counterpartyCh.ID,
				},
				ConnectionHops: []string{connectionID},
				Version:        ch.Version,
			},
			CounterpartyVersion: counterpartyCh.Version,
			ProofInit:           proof.Data,
			ProofHeight:         proof.Height.ToCallData(),
		},
	); err != nil {
		return "", err
	}
	return chain.GetLastGeneratedChannelID(ctx)
}

func (chain *Chain) ChannelOpenAck(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelOpenAck",
		ibchandler.IBCMsgsMsgChannelOpenAck{
			PortId:                ch.PortID,
			ChannelId:             ch.ID,
			CounterpartyVersion:   counterpartyCh.Version,
			CounterpartyChannelId: counterpartyCh.ID,
			ProofTry:              proof.Data,
			ProofHeight:           proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) ChannelOpenConfirm(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelOpenConfirm",
		ibchandler.IBCMsgsMsgChannelOpenConfirm{
			PortId:      ch.PortID,
			ChannelId:   ch.ID,
			ProofAck:    proof.Data,
			ProofHeight: proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) ChannelCloseInit(
	ctx context.Context,
	ch TestChannel,
) error {
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelCloseInit",
		ibchandler.IBCMsgsMsgChannelCloseInit{
			PortId:    ch.PortID,
			ChannelId: ch.ID,
		},
	)
}

func (chain *Chain) ChannelCloseConfirm(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
) error {
	proof, err := counterparty.QueryChannelProof(chain, ch.ClientID, counterpartyCh, nil)
	if err != nil {
		return err
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"channelCloseConfirm",
		ibchandler.IBCMsgsMsgChannelCloseConfirm{
			PortId:      ch.PortID,
			ChannelId:   ch.ID,
			ProofInit:   proof.Data,
			ProofHeight: proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) SendPacket(
	ctx context.Context,
	packet channeltypes.Packet,
) error {
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"sendPacket",
		packetToCallData(packet),
	)
}

func (chain *Chain) HandlePacketRecv(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet channeltypes.Packet,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.PacketCommitmentSlot(packet.SourcePort, packet.SourceChannel, packet.Sequence), nil)
	if err != nil {
		return err
	}
	switch chain.ClientType() {
	case ibcclient.MockClient:
		proof.Data = commitPacket(packet)
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"recvPacket",
		ibchandler.IBCMsgsMsgPacketRecv{
			Packet:      packetToCallData(packet),
			Proof:       proof.Data,
			ProofHeight: proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) HandlePacketAcknowledgement(
	ctx context.Context,
	counterparty *Chain,
	ch, counterpartyCh TestChannel,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	proof, err := counterparty.QueryProof(chain, ch.ClientID, chain.PacketAcknowledgementCommitmentSlot(packet.DestinationPort, packet.DestinationChannel, packet.Sequence), nil)
	if err != nil {
		return err
	}
	switch chain.ClientType() {
	case ibcclient.MockClient:
		proof.Data = commitAcknowledgement(acknowledgement)
	}
	return chain.transactAndWaitMined(
		chain.TxOpts(ctx, RelayerKeyIndex),
		"acknowledgePacket",
		ibchandler.IBCMsgsMsgPacketAcknowledgement{
			Packet:          packetToCallData(packet),
			Acknowledgement: acknowledgement,
			Proof:           proof.Data,
			ProofHeight:     proof.Height.ToCallData(),
		},
	)
}

func (chain *Chain) GetLastGeneratedClientID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedClientIdentifier)
}

func (chain *Chain) GetLastGeneratedConnectionID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedConnectionIdentifier)
}

func (chain *Chain) GetLastGeneratedChannelID(
	ctx context.Context,
) (string, error) {
	return chain.getLastID(ctx, abiGeneratedChannelIdentifier)
}

func (chain *Chain) getLastID(ctx context.Context, event abi.Event) (string, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.ContractConfig.GetIBCHostAddress(),
		},
		Topics: [][]common.Hash{{
			event.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx, query)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", errors.New("no items")
	}
	log := logs[len(logs)-1]
	values, err := event.Inputs.Unpack(log.Data)
	if err != nil {
		return "", err
	}
	return values[0].(string), nil
}

func (chain *Chain) GetLastSentPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
) (*channeltypes.Packet, error) {
	seq, err := chain.IBCHost.GetNextSequenceSend(chain.CallOpts(ctx, RelayerKeyIndex), sourcePortID, sourceChannel)
	if err != nil {
		return nil, err
	}
	return chain.FindPacket(ctx, sourcePortID, sourceChannel, seq-1)
}

func (chain *Chain) FindPacket(
	ctx context.Context,
	sourcePortID string,
	sourceChannel string,
	sequence uint64,
) (*channeltypes.Packet, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			chain.ContractConfig.GetIBCHandlerAddress(),
		},
		Topics: [][]common.Hash{{
			abiSendPacket.ID,
		}},
	}
	logs, err := chain.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, log := range logs {
		if values, err := abiSendPacket.Inputs.Unpack(log.Data); err != nil {
			return nil, err
		} else {
			p := values[0].(struct {
				Sequence           uint64  "json:\"sequence\""
				SourcePort         string  "json:\"source_port\""
				SourceChannel      string  "json:\"source_channel\""
				DestinationPort    string  "json:\"destination_port\""
				DestinationChannel string  "json:\"destination_channel\""
				Data               []uint8 "json:\"data\""
				TimeoutHeight      struct {
					RevisionNumber uint64 "json:\"revision_number\""
					RevisionHeight uint64 "json:\"revision_height\""
				} "json:\"timeout_height\""
				TimeoutTimestamp uint64 "json:\"timeout_timestamp\""
			})
			if p.SourcePort == sourcePortID && p.SourceChannel == sourceChannel && p.Sequence == sequence {
				return &channeltypes.Packet{
					Sequence:           p.Sequence,
					SourcePort:         p.SourcePort,
					SourceChannel:      p.SourceChannel,
					DestinationPort:    p.DestinationPort,
					DestinationChannel: p.DestinationChannel,
					Data:               p.Data,
					TimeoutHeight:      ibcclient.Height(p.TimeoutHeight),
					TimeoutTimestamp:   p.TimeoutTimestamp,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("packet not found: sourcePortID=%v sourceChannel=%v sequence=%v", sourcePortID, sourceChannel, sequence)
}

func packetToCallData(packet channeltypes.Packet) ibchandler.PacketData {
	return ibchandler.PacketData{
		Sequence:           packet.Sequence,
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               packet.Data,
		TimeoutHeight:      ibchandler.HeightData(packet.TimeoutHeight),
		TimeoutTimestamp:   packet.TimeoutTimestamp,
	}
}

// Slot calculator

func (chain *Chain) ClientStateCommitmentSlot(clientID string) string {
	key, err := chain.IBCIdentifier.ClientStateCommitmentSlot(chain.CallOpts(context.Background(), RelayerKeyIndex), clientID)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) ConnectionStateCommitmentSlot(connectionID string) string {
	key, err := chain.IBCIdentifier.ConnectionCommitmentSlot(chain.CallOpts(context.Background(), RelayerKeyIndex), connectionID)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) ChannelStateCommitmentSlot(portID, channelID string) string {
	key, err := chain.IBCIdentifier.ChannelCommitmentSlot(chain.CallOpts(context.Background(), RelayerKeyIndex), portID, channelID)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) PacketCommitmentSlot(portID, channelID string, sequence uint64) string {
	key, err := chain.IBCIdentifier.PacketCommitmentSlot(chain.CallOpts(context.Background(), RelayerKeyIndex), portID, channelID, sequence)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

func (chain *Chain) PacketAcknowledgementCommitmentSlot(portID, channelID string, sequence uint64) string {
	key, err := chain.IBCIdentifier.PacketAcknowledgementCommitmentSlot(chain.CallOpts(context.Background(), RelayerKeyIndex), portID, channelID, sequence)
	require.NoError(chain.t, err)
	return "0x" + hex.EncodeToString(key[:])
}

// Querier

type Proof struct {
	Height *ibcclient.Height
	Data   []byte
}

func (chain *Chain) QueryProof(counterparty *Chain, counterpartyClientID string, storageKey string, height *big.Int) (*Proof, error) {
	if !strings.HasPrefix(storageKey, "0x") {
		return nil, fmt.Errorf("storageKey must be hex string")
	}
	s, err := chain.GetContractState(counterparty, counterpartyClientID, [][]byte{[]byte(storageKey)}, height)
	if err != nil {
		return nil, err
	}
	return &Proof{
		Height: ibcclient.NewHeightFromBN(s.Header().Number),
		Data:   s.ETHProof().StorageProofRLP[0],
	}, nil
}

func (counterparty *Chain) QueryClientProof(chain *Chain, counterpartyClientID string, height *big.Int) ([]byte, *Proof, error) {
	cs, found, err := counterparty.IBCHost.GetClientState(
		counterparty.CallOpts(context.Background(), RelayerKeyIndex),
		counterpartyClientID,
	)
	if err != nil {
		return nil, nil, err
	} else if !found {
		return nil, nil, fmt.Errorf("client not found: %v", counterpartyClientID)
	}
	proof, err := counterparty.QueryProof(chain, counterpartyClientID, chain.ClientStateCommitmentSlot(counterpartyClientID), height)
	if err != nil {
		return nil, nil, err
	}
	switch counterparty.ClientType() {
	case ibcclient.MockClient:
		h := sha256.Sum256(cs)
		proof.Data = h[:]
	}
	return cs, proof, nil
}

func (counterparty *Chain) QueryConnectionProof(chain *Chain, counterpartyClientID string, counterpartyConnectionID string, height *big.Int) (*Proof, error) {
	proof, err := counterparty.QueryProof(chain, counterpartyClientID, chain.ConnectionStateCommitmentSlot(counterpartyConnectionID), height)
	if err != nil {
		return nil, err
	}
	switch counterparty.ClientType() {
	case ibcclient.MockClient:
		conn, found, err := counterparty.IBCHost.GetConnection(
			counterparty.CallOpts(context.Background(), RelayerKeyIndex),
			counterpartyConnectionID,
		)
		if err != nil {
			return nil, err
		} else if !found {
			return nil, fmt.Errorf("connection not found: %v", counterpartyConnectionID)
		}
		bz, err := proto.Marshal(connectionEndToPB(conn))
		if err != nil {
			return nil, err
		}
		h := sha256.Sum256(bz)
		proof.Data = h[:]
	}
	return proof, nil
}

func (counterparty *Chain) QueryChannelProof(chain *Chain, counterpartyClientID string, channel TestChannel, height *big.Int) (*Proof, error) {
	proof, err := counterparty.QueryProof(chain, counterpartyClientID, chain.ChannelStateCommitmentSlot(channel.PortID, channel.ID), height)
	if err != nil {
		return nil, err
	}
	switch counterparty.ClientType() {
	case ibcclient.MockClient:
		ch, found, err := counterparty.IBCHost.GetChannel(
			counterparty.CallOpts(context.Background(), RelayerKeyIndex),
			channel.PortID, channel.ID,
		)
		if err != nil {
			return nil, err
		} else if !found {
			return nil, fmt.Errorf("channel not found: %v", channel)
		}
		bz, err := proto.Marshal(channelToPB(ch))
		if err != nil {
			return nil, err
		}
		h := sha256.Sum256(bz)
		proof.Data = h[:]
	}
	return proof, nil
}

func (chain *Chain) LastHeader() *gethtypes.Header {
	return chain.LastContractState.Header()
}

func (chain *Chain) WaitForReceiptAndGet(ctx context.Context, tx *gethtypes.Transaction) error {
	rc, err := chain.Client().WaitForReceiptAndGet(ctx, tx)
	if err != nil {
		return err
	}
	if rc.Status() == 1 {
		return nil
	} else {
		return fmt.Errorf("failed to call transaction: err='%v' rc='%v' reason='%v'", err, rc, rc.RevertReason())
	}
}

func (chain *Chain) WaitIfNoError(ctx context.Context) func(tx *gethtypes.Transaction, err error) error {
	return func(tx *gethtypes.Transaction, err error) error {
		if err != nil {
			return err
		}
		if err := chain.WaitForReceiptAndGet(ctx, tx); err != nil {
			return err
		}
		return nil
	}
}

// AddTestConnection appends a new TestConnection which contains references
// to the connection id, client id and counterparty client id.
func (chain *Chain) AddTestConnection(clientID, counterpartyClientID string) *TestConnection {
	conn := chain.ConstructNextTestConnection(clientID, counterpartyClientID)

	chain.Connections = append(chain.Connections, conn)
	return conn
}

// ConstructNextTestConnection constructs the next test connection to be
// created given a clientID and counterparty clientID.
func (chain *Chain) ConstructNextTestConnection(clientID, counterpartyClientID string) *TestConnection {
	return &TestConnection{
		ID:                   "",
		ClientID:             clientID,
		NextChannelVersion:   DefaultChannelVersion,
		CounterpartyClientID: counterpartyClientID,
	}
}

// AddTestChannel appends a new TestChannel which contains references to the port and channel ID
// used for channel creation and interaction. See 'NextTestChannel' for channel ID naming format.
func (chain *Chain) AddTestChannel(conn *TestConnection, portID string) TestChannel {
	channel := chain.NextTestChannel(conn, portID)
	conn.Channels = append(conn.Channels, channel)
	return channel
}

// NextTestChannel returns the next test channel to be created on this connection, but does not
// add it to the list of created channels. This function is expected to be used when the caller
// has not created the associated channel in app state, but would still like to refer to the
// non-existent channel usually to test for its non-existence.
//
// The port is passed in by the caller.
func (chain *Chain) NextTestChannel(conn *TestConnection, portID string) TestChannel {
	return TestChannel{
		PortID:               portID,
		ID:                   "",
		ClientID:             conn.ClientID,
		CounterpartyClientID: conn.CounterpartyClientID,
		Version:              conn.NextChannelVersion,
	}
}
