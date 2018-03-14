package blockchain

// import (
// 	"crypto/rsa"
// 	"encoding/pem"
// )

// func SendKey(key *rsa.PrivateKey) pem.Block {
// 	var privateKey = &pem.Block{
// 		Type:  "PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(key),
// 	}

// 	err = pem.Encode(outFile, privateKey)

// }

// func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
// 	asn1Bytes, err := asn1.Marshal(pubkey)
// 	checkError(err)

// 	var pemkey = &pem.Block{
// 		Type:  "PUBLIC KEY",
// 		Bytes: asn1Bytes,
// 	}

// 	pemfile, err := os.Create(fileName)

// 	defer pemfile.Close()

// 	err = pem.Encode(pemfile, pemkey)
// }
