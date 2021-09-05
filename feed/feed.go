package feed

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/lightclient/rols/feed/bindings"
)

type Backend interface {
	bind.ContractBackend

	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
	SubscribeNewHead(context.Context, chan<- *types.Header) (ethereum.Subscription, error)
}

type Config struct {
	ProviderUrl string
	FeedAddress common.Address
}

type Feed struct {
	ctx      context.Context
	stop     chan struct{}
	contract *bindings.CanonicalTransactionChainFilterer
	backend  Backend
}

func New(cfg *Config) (*Feed, error) {
	client, err := ethclient.Dial(cfg.ProviderUrl)
	if err != nil {
		return nil, err
	}

	address := cfg.FeedAddress
	contract, err := bindings.NewCanonicalTransactionChainFilterer(address, client)
	if err != nil {
		return nil, err
	}

	return &Feed{
		ctx:      context.Background(),
		stop:     make(chan struct{}),
		contract: contract,
		backend:  client,
	}, nil
}

func (f *Feed) Start() {
	newHeads := make(chan *types.Header, 1000)

	go func() {
		subscription, err := f.backend.SubscribeNewHead(f.ctx, newHeads)
		if err != nil {
			panic(fmt.Sprintf("Unable to subscribe to new heads: %v", err))
		}
		defer subscription.Unsubscribe()
		for {
			select {
			case header := <-newHeads:
				log.Info("Received new header", header)
			case <-f.stop:
				return
			}
		}
	}()
}
