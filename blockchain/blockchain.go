package blockchain

import (
	"crypto/sha256"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type BlockChain struct {
	Index               int
	CurrentTransactions []Transaction
	Chain               []Block
}

func (self BlockChain) GenerateUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	// fmt.Printf("UUIDv4: %s\n", u1)
	return u1.String()

}

// Creates a new Block and adds it to the chain
// proof: The proof given by the Proof of Work algorithm
// previous_hash: (Optional) <str> Hash of previous Block
func (self *BlockChain) NewBlock(proof int, previous_hash string) {

	// check if previous hash matches self.hash(self.chain[-1])
	t := time.Now()

	block := Block{
		Index:        len(self.Chain) + 1,
		Timestamp:    t.UnixNano(),
		Transactions: self.CurrentTransactions,
		Proof:        proof,
		PreviousHash: previous_hash}

	// Reset the current list of transactions
	self.CurrentTransactions = nil
	self.Chain = append(self.Chain, block)
}

//  Adds a new transaction to the list of transactions
func (self *BlockChain) NewTransaction(sender string, recipient string, amount float32) {
	newTransaction := Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	self.CurrentTransactions = append(self.CurrentTransactions, newTransaction)
	print("added. pending transactions ... ")
	print(len(self.CurrentTransactions), "\n")

}

//Debugging Helper function
func (self BlockChain) PrintInfo() {
	fmt.Printf("Pending Transactions  - %d  \n", len(self.CurrentTransactions))
	fmt.Printf("Length of Blockchain %d \n", len(self.Chain))
	fmt.Printf("Blockchain data %v", (self.Chain))
}

// Simple Proof of Work Algorithm:
// - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
// - p is the previous proof, and p' is the new proof
func (self *BlockChain) ProofOfWork(last_proof int) int {
	proof := 0
	for self.ValidProof(last_proof, proof) == false {
		proof += 1
	}
	println("Found the proof!!")
	println(proof)
	return proof
}

// Validates the Proof: Does hash(last_proof, proof) contain 4 leading zeroes?
func (self *BlockChain) ValidProof(lastProof int, proof int) bool {
	// guess = f'{last_proof}{proof}'.encode()
	combined := fmt.Sprintf("%d%d", lastProof, proof)
	sum := sha256.Sum256([]byte(combined))
	hex := fmt.Sprintf("%x", sum)

	// guess_hash := hashlib.sha256(guess).hexdigest()
	// println(hex)
	return hex[len(hex)-4:] == "0000"
}

//Mine a coin rewarding the user with a coin, and create a block.
func (self *BlockChain) Mine() {
	//Get the last transactions proof
	println("Starting to mine...")
	last := self.LastBlock()
	lastProof := last.Proof

	//Work out the proof
	fmt.Printf("Last proof  =  %d \n", lastProof)
	newProof := self.ProofOfWork(lastProof)
	self.NewTransaction("0", "dest", 1)

	//Add to blockchain with Proof + HASH
	//TODO get hash of previous
	fmt.Printf("new proof  =  %d \n", newProof)
	self.NewBlock(newProof, "xxx")

}

// # Hashes a Block, static method, pass in something else?
func (b BlockChain) Hash() {

}

// Returns the last Block in the chain
func (self *BlockChain) LastBlock() Block {
	println("getting last block")
	return self.Chain[len(self.Chain)-1]
}
