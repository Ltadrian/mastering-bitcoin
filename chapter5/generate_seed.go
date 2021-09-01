package chapter5

import (
	"crypto/sha512"
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
