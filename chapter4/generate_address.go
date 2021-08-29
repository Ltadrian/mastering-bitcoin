package chapter4

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/anaskhan96/base58check"
	"golang.org/x/crypto/ripemd160"
)

func GenerateKey() *ecdsa.PrivateKey {
	var secret, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate key: %v\n", err)
		os.Exit(1)
	}
	return secret
}

func GenerateBitcoinAddress(secret *ecdsa.PrivateKey) string {
	encodedPublicKeyX := hex.EncodeToString(secret.PublicKey.X.Bytes())
	encodedPublicKeyY := hex.EncodeToString(secret.PublicKey.Y.Bytes())

	fmt.Println("Public Key - X Coordinate: " + encodedPublicKeyX)
	fmt.Println("Public Key - Y Coordinate: " + encodedPublicKeyY)

	uncompressedPublicKey := "04" + encodedPublicKeyX + encodedPublicKeyY
	fmt.Println("Uncompressed Public Key: " + uncompressedPublicKey)

	var compressedKey string
	if secret.PublicKey.Y.Bit(0) == 0 { // big int even check
		compressedKey = "02" + encodedPublicKeyX
	} else { // otherwise odd
		compressedKey = "03" + encodedPublicKeyX
	}
	fmt.Println("Compressed Public Key: " + compressedKey)

	// Take Compressed Public Key and generate a valid Bitcoin Address
	// A = RIPEMD160(SHA256(K))
	sum256Encoded := sha256.Sum256([]byte(compressedKey)) // SHA256(K)
	hasher := ripemd160.New()                             // RIPEMD160
	hasher.Write(sum256Encoded[:])                        // convert [32] byte to []byte by creating slice
	hashBytes := hasher.Sum(nil)
	A := fmt.Sprintf("%x", hashBytes)
	fmt.Println("Address before Base58Check encoding: " + A)

	// Base58 Check Encoding
	base58EncodedAddress, err := base58check.Encode("00", hex.EncodeToString(hashBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "base58check: %v\n", err)
	}
	fmt.Printf("Successfully generated Base58Check encoded bitcoin address: %s\n", base58EncodedAddress)
	return base58EncodedAddress
}
