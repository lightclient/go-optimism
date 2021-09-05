package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// A cross-domain message on the canonical chain.
type Message struct {
	QueueIndex    uint64
	CanonTxHash   common.Hash
	CanonTxOrigin common.Address
	Target        common.Address
	GasLimit      uint64
	Data          []byte
}
