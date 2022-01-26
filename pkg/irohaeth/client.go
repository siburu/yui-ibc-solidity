package irohaeth

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	c *rpc.Client
}

func Dial(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

func (c *Client) Close() {
	c.c.Close()
}

type Transaction struct {
	*types.Transaction
	ID common.Hash
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) (*Transaction, error) {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, err
	}
	var txID common.Hash
	err = c.c.CallContext(ctx, &txID, "eth_sendRawTransaction", hexutil.Encode(data))
	if err != nil {
		return nil, err
	}
	return &Transaction{
		Transaction: tx,
		ID:          txID,
	}, nil
}

func (c *Client) TransactionReceipt(ctx context.Context, txID common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := c.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txID)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

func (c *Client) WaitMined(ctx context.Context, tx *Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := c.TransactionReceipt(ctx, tx.ID)
		if err != nil {
			return nil, err
		} else if receipt != nil {
			return receipt, nil
		} else {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		}
	}
}
