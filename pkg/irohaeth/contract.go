package irohaeth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BoundContract struct {
	address common.Address
	abi     abi.ABI
	client  *Client
}

func NewBoundContract(address common.Address, abi abi.ABI, client *Client) *BoundContract {
	return &BoundContract{
		address: address,
		abi:     abi,
		client:  client,
	}
}

func (c *BoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*Transaction, error) {
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	nonce := uint64(0)
	amount := big.NewInt(0)
	gasLimit := uint64(0)
	gasPrice := big.NewInt(0)

	// make a transaction
	rawTx := types.NewTransaction(nonce, c.address, amount, gasLimit, gasPrice, input)
	if opts.Signer == nil {
		panic("no signer to authorize the transaction with")
	}

	// sign the transaction
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	// send the transaction
	ctx := opts.Context
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.client.SendTransaction(ctx, signedTx)
}
