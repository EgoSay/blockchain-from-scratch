package tests

import (
	"blockchain-from-scratch/core"
	"fmt"
	"testing"
)

func TestInitChain(t *testing.T) {
	chain := core.NewBlockChain()

	//chain.MineBlock("Send tx1")
	//chain.MineBlock("Send tx2")

	iterator := chain.Iterator()
	for {
		block := iterator.Next()
		if len(block.PrevBlockHash) == 0 {
			break
		}
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.HashTransactions())
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
