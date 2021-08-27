package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	var secret, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate key: %v\n", err)
	}

	encodedPublicKeyX := hex.EncodeToString(secret.PublicKey.X.Bytes())
	encodedPublicKeyY := hex.EncodeToString(secret.PublicKey.Y.Bytes())

	fmt.Println("x:" + encodedPublicKeyX)
	fmt.Println("y:" + encodedPublicKeyY)

	uncompressedPublicKey := "04" + encodedPublicKeyX + encodedPublicKeyY
	fmt.Println("uncompressed public key:" + uncompressedPublicKey)

	var compressedKey string
	if secret.PublicKey.Y.Bit(0) == 0 { // big int even check
		compressedKey = "02" + encodedPublicKeyX
	} else { // otherwise odd
		compressedKey = "03" + encodedPublicKeyX
	}
	fmt.Println("compressed public key:" + compressedKey)

	sum256Encoded := sha256.Sum256([]byte(compressedKey))
	hasher := ripemd160.New()
	hasher.Write(sum256Encoded[:]) // convert [32] byte to []byte by creating slice
	hashBytes := hasher.Sum(nil)
	A := fmt.Sprintf("%x", hashBytes)
	fmt.Println("address before base58 encoding:" + A)

	base58EncodedAddress := base58.Encode(hashBytes)
	fmt.Printf("Successfully base58 encoded address: %s\n", base58EncodedAddress)
}
