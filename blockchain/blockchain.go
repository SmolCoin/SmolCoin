package blockchain

import (
	"time"

	"github.com/SmolCoin/SmolCoin/block"
)

type Chain []block.Block

func New() Chain {
	return Chain{block.Genisis()}
}

func (chain *Chain) Dump() {
	for _, block := range *chain {
		block.Dump()
	}
}

// TODO: When we have data to put in this is going to be an argument
func (chain *Chain) PushBlock() block.Block {
	the := block.New((*chain)[len(*chain)-1].Hash, time.Now())
	*chain = append(*chain, the)

	return the
}
