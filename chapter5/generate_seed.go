package chapter5

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

/**
  Generate 512-bit seed from Mnemonic Code Words (BIP-39)
*/
func GenerateSeed(wordList []string, salt string) []byte {
	// PBKDF2 using HMAC-SHA512 with 2048 rounds to produce 512-bit Seed
	return pbkdf2.Key(
		[]byte(strings.Join(wordList, " ")),
		[]byte(salt),
		2048,
		64,
		sha512.New,
	)
}

func GenerateMasterKeys(seed []byte) (masterPrivateKey, masterChainCode string) {
	sha512hmac := hmac.New(sha512.New, []byte{})
	_, err := sha512hmac.Write(seed)
	if err != nil {
		panic(err)
	}

	// Get result and encode as hexadecimal string
	sha512hmac_b := sha512hmac.Sum(nil)
	masterPrivateKey = hex.EncodeToString(sha512hmac_b[:32]) // left 256 bits
	masterChainCode = hex.EncodeToString(sha512hmac_b[32:])  // right 256 bits
	return masterPrivateKey, masterChainCode
}
