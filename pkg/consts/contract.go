package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0xd5d497E3e3b2589DDB3eaE9b644c70218A991826"
	IBCHandlerAddress = "0x01F8409515Ed5A5909ABcb25058e7fE6D0727ee3"
	IBCIdentifierAddress = "0x13aA93254b9F267FCFa453b38f69D15960FEd9d4"
	IBFT2ClientAddress = "0x1Ba5ca2AEf1B4A3a3bc0420Cc354dad4aAe40EA4"
	MockClientAddress = "0x2BC5611386919fa2F9a1eeBeb2A7fd2eC22d00Fe"
	SimpleTokenAddress = "0x3fFA9d74F012C56869A978aD80e76d5F39D93D18"
	ICS20TransferBankAddress = "0xC9ead3f21Ba7F3D87b432cF4E1a2fDe8bE1bfD5E"
	ICS20BankAddress = "0xEa21C99b6AeDa44737301F6433A61ED070684157"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetMockClientAddress() common.Address {
	return common.HexToAddress(MockClientAddress)
}

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferBankAddress() common.Address {
	return common.HexToAddress(ICS20TransferBankAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
