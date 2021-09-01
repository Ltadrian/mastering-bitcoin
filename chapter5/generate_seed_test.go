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
