package main

import (
	"encoding/json"
	"fmt"
	"github.com/bomer/blockchain/blockchain"
	"net/http"
)

var (
	MyBlockChain blockchain.BlockChain
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there. Seems you hit a dummy end point! Bumblebay tuna!")
}

//Add a transaction, need the following post params
// Receives json obj of type Block
func addhandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var data blockchain.Transaction
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	MyBlockChain.NewTransaction(data.Recipient, data.Recipient, data.Amount)
	json.NewEncoder(w).Encode(MyBlockChain.CurrentTransactions)

}
func minehandler(w http.ResponseWriter, r *http.Request) {
	MyBlockChain.Mine()
	json.NewEncoder(w).Encode(MyBlockChain.Chain)
}
func chainhandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MyBlockChain.Chain)
}
func transactionshandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MyBlockChain.CurrentTransactions)
}

func main() {
	//SETUP the genisys
	MyBlockChain.NewBlock(100, "1") // Genysis Block

	//HTTP Routes
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/", handler)
	http.HandleFunc("/api/add", addhandler)
	http.HandleFunc("/api/mine", minehandler)
	http.HandleFunc("/api/chain", chainhandler)
	http.HandleFunc("/api/transactions", transactionshandler)

	http.ListenAndServe(":8008", nil)
}
