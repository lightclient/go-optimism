package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO: does this even need to be stored?
//
// Snippets of a block from the canonical chain, including messages to the
// child chain and batches of child chain transactions.
type CanonBlock struct {
	ParentHash common.Hash
	Root       common.Hash
	Number     uint64
	Time       uint64
	Messages   []*Message
	Batches    [][]*types.Transaction
}
