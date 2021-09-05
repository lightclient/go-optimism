package node

import (
	log "github.com/inconshreveable/log15"
	"github.com/lightclient/rols/feed"
)

type Config struct {
	// URL of feed oracle.
	FeedOracle string

	//
	FeedContract string

	// URL of sequencer RPC.
	Sequencer string

	// Private key for transaction signer.
	SigningKey string

	// Number of blocks to trial L1.
	ConfirmationDepth uint64
}

type Node struct {
	feed *feed.Feed
}

func New(conf *Config) (*Node, error) {
	return &Node{}, nil
}

func (n *Node) Start() error {
	log.Info("Starting rollup synchronizer")
	return nil
}
