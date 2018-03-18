package blockchain

// Transaction to be aded to a block
type Transaction struct {
	Sender    string
	Recipient string
	Amount    float32
	// Hash      string
}

//Blocks to be added to the chain
type Block struct {
	Index        int
	Timestamp    int64
	Proof        int
	PreviousHash string
	Transactions []Transaction
}
