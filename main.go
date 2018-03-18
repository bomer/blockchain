package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	// "encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/bomer/blockchain/blockchain"
	"math/big"
	"net/http"
	"strconv"
)

var (
	MyBlockChain blockchain.BlockChain
)

type key struct {
	alg string
	E   string
	ext bool
	kty string
	N   string
}
type addrequest struct {
	Signature   map[string]uint8
	Key         key
	Transaction blockchain.Transaction //map[string]int //
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there. Seems you hit a dummy end point! Bumblebay tuna!")
}

// type JSONableSlice []uint8

// func (u JSONableSlice) MarshalJSON() ([]byte, error) {
// 	var result string
// 	if u == nil {
// 		result = "null"
// 	} else {
// 		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
// 	}
// 	return []byte(result), nil
// }

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

	//Extract the E + N (exponent + Number) of the encoded public Key.
	z := new(big.Int)
	byteslice := []byte(data.Key.N)
	z.SetBytes(byteslice)

	e_data, _ := base64.RawURLEncoding.DecodeString(data.Key.E)

	e_big := new(big.Int)
	e_big.SetBytes(e_data)
	e := int(e_big.Int64())

	readPublicKey := rsa.PublicKey{N: z, E: e}

	fmt.Printf("Publick key ========== \n %v \n", readPublicKey)
	// print(crypto.SHA256)

	// bytemsg := []byte(data.Transaction)
	// print(data.Transaction)
	b, err := json.Marshal(data.Transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(b))
	// bytesOfTransaction := []byte(b)
	// fmt.Printf("transaction %v \n", bytesOfTransaction)
	shaOfBytesOfTransaction := sha256.Sum256(b)
	// str := hex.EncodeToString(shaOfBytesOfTransaction[:])
	// hashOfBytesOfTransaction := []byte(str)
	// test := string(hashOfBytesOfTransaction)
	hashOfBytesOfTransaction := shaOfBytesOfTransaction[:]

	// fmt.Printf("hased trans string == \n%v\n", str)
	fmt.Printf("hased trans == bytes \n%v\n", hashOfBytesOfTransaction)
	// var v []byte
	// for _, value := range data.Signature {
	// 	v = append(v, byte(value))
	// }
	v := data.Signature
	// fmt.Printf("Signature %v", v)
	var signatureByteSlice []byte

	for i := 0; i < len(v); i++ {
		s := strconv.Itoa(i)
		out := v[s]
		// fmt.Printf("%d - %d \n", i, out)
		signatureByteSlice = append(signatureByteSlice, byte(out))

	}
	fmt.Printf("sig byte slice trans == bytes \n %v \n", signatureByteSlice)

	// sig := new([]byte)
	validationerror := rsa.VerifyPKCS1v15(&readPublicKey, crypto.SHA256, shaOfBytesOfTransaction[:], signatureByteSlice)
	if validationerror != nil {
		fmt.Printf("Error from verification: %s\n", validationerror)
		return
	}

	// MyBlockChain.NewTransaction(data.Transaction.Sender, data.Transaction.Recipient, data.Transaction.Amount)
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
