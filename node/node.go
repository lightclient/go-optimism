package node

import (
	log "github.com/inconshreveable/log15"
	"github.com/lightclient/go-optimism/feed"
)

type Config struct {
	// URL of feed oracle.
	FeedProvider string

	// Canon address of feed oracle.
	FeedContract string
}

type Node struct {
	feed *feed.Feed
}

func New(conf *Config) (*Node, error) {
	return &Node{}, nil
}

func (n *Node) Start() error {
	log.Info("Starting go-optimism . . .")
	return nil
}
