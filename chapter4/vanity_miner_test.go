package chapter4

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMatchFound(t *testing.T) {
	address := "1kidLJfdGaP4EeHnDyJbEGQysnCpwk3gd"
	vanityKey := "1kid"

	validVanityAddress := MatchFound(address, vanityKey)

	assert.Equal(t, true, validVanityAddress)
}

func TestMatchFoundCaseInsensitive(t *testing.T) {
	address := "1kidLJfdGaP4EeHnDyJbEGQysnCpwk3gd"
	vanityKey := "1KID"

	validVanityAddress := MatchFound(address, vanityKey)

	assert.Equal(t, true, validVanityAddress)
}

func TestMatchFoundNotValid(t *testing.T) {
	address := "14qViLJfdGaP4EeHnDyJbEGQysnCpwk3gd"
	vanityKey := "1kid"

	validVanityAddress := MatchFound(address, vanityKey)

	assert.Equal(t, false, validVanityAddress)
}
