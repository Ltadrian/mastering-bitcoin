package chapter5

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenerateSeed(t *testing.T) {
	masterSeed := GenerateSeed(English[:12], "mnemonic")
	fmt.Printf("master seed:%s\n len:%d\n", hex.EncodeToString(masterSeed), len(masterSeed))
	assert.Equal(t, 64, len(masterSeed)) // 512 bits === 64 bytes
}

func TestGenerateMasterKeys(t *testing.T) {
	masterPrivateKey, masterChainCode := GenerateMasterKeys(GenerateSeed(English[:12], "mnemonic"))
	fmt.Printf("masterPrivateKey:%s\nlen:%d\n", masterPrivateKey, len(masterPrivateKey))
	fmt.Printf("masterChainCode:%s\nlen:%d\n", masterChainCode, len(masterChainCode))
	assert.Equal(t, 64, len(masterPrivateKey))
	assert.Equal(t, 64, len(masterChainCode))
}
