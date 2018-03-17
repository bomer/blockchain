package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/bomer/blockchain/blockchain"
	"net/http"
)

var (
	MyBlockChain blockchain.BlockChain
)

type addrequest struct {
	Data        string
	Transaction blockchain.Transaction
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there. Seems you hit a dummy end point! Bumblebay tuna!")
}

//Add a transaction, need the following post params
// Receives json obj of type Block
func addhandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var data addrequest
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	MyBlockChain.NewTransaction(data.Transaction.Sender, data.Transaction.Recipient, data.Transaction.Amount)
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

func generatekeyhandler(w http.ResponseWriter, r *http.Request) {

	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey
	// json.NewEncoder(w).Encode(publicKey)
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	fmt.Fprintf(w, "{\"private\":\"")
	_ = pem.Encode(w, privateKey)
	fmt.Fprintf(w, "\",\"public\":\"")
	//now pub;ic key
	asn1Bytes, err := asn1.Marshal(publicKey)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	checkError(err)
	_ = pem.Encode(w, pemkey)
	fmt.Fprintf(w, "\"}")

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
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
	http.HandleFunc("/api/generatekeys", generatekeyhandler)

	http.ListenAndServe(":8008", nil)
}
