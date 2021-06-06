package main

import (
	"fmt"

	"github.com/SmolCoin/SmolCoin/blockchain"
)

func main() {
	chain := blockchain.New()
	chain.Dump()

	fmt.Println("-------------------------")
	chain.PushBlock()
	chain.PushBlock()
	chain.Dump()
}
