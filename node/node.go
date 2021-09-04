package node

import log "github.com/inconshreveable/log15"

type Node struct {
}

type Config struct {
	FeedOracle string // URL of feed oracle RPC
	Sequencer  string // URL of sequencer RPC
	SigningKey string // private key for transaction signer
}

func New(conf *Config) (*Node, error) {
	return &Node{}, nil
}

func (n *Node) Start() error {
	log.Info("Starting rollup synchronizer")
	return nil
}
