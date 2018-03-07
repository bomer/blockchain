package main

import (
	"fmt"
	"github.com/bomer/blockchain/blockchain"
)

var (
	MyBlockChain blockchain.BlockChain
)

func main() {
	println("hi")

	testUUID := MyBlockChain.GenerateUUID()
	fmt.Printf("test UUID: %s\n", testUUID)

	MyBlockChain.Index = 1
	MyBlockChain.NewBlock(100, "1") // Genysis Block

	MyBlockChain.PrintInfo()
	MyBlockChain.NewTransaction("123", "234", 1.0)
	MyBlockChain.NewTransaction("123", "234", 5.0)
	MyBlockChain.NewTransaction("234", "123", 2.0)
	MyBlockChain.PrintInfo()

	// println("Now adding my first block... \n")
	// MyBlockChain.NewBlock(1, "999")
	MyBlockChain.Mine()
	println()
	MyBlockChain.PrintInfo()

	println("Trying Proof of Work \n")
	// MyBlockChain.ProofOfWork(232)
	MyBlockChain.NewTransaction("234", "123", 2.0)
	MyBlockChain.Mine()
	// MyBlockChain.NewBlock(1, "999")

	MyBlockChain.NewTransaction("234", "123", 2.0)
	MyBlockChain.PrintInfo()

	println("Now adding my 2nd block...  \n")
	MyBlockChain.Mine()
	MyBlockChain.PrintInfo()

	MyBlockChain.NewTransaction("234", "123", 2.0)
	MyBlockChain.PrintInfo()

	println("Now adding my 4th block...  \n")
	MyBlockChain.Mine()
	MyBlockChain.PrintInfo()
}
