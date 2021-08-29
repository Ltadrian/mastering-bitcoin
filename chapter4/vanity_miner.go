package chapter4

import (
	"fmt"
	"strings"
	"time"
)

/* Average Search Times
	length | pattern |	frequency  		| average search time
	1        1K			1 in 58				< 1 ms
	2		 1Ki		1 in 3,364			50 ms
	3		 1Kid		1 in 195,000		< 2 seconds
  	4		 1Kids		1 in 11 million		1 minute
	5		 1KidsC		1 in 656 million	1 hour
	...
	11 | 1KidsCharity | 1 in 23 quintillion |	2.5 million years
*/
func MineVanityAddress(vanityKey string) {
	startTime := time.Now()
	for {
		secret := GenerateKey()
		address := GenerateBitcoinAddress(secret)
		if MatchFound(address, vanityKey) {
			currentTime := time.Now()
			diff := currentTime.Sub(startTime)
			fmt.Printf("Vanity address found: %s\n", address)
			fmt.Printf("Vanity Key: %s\n", vanityKey)
			fmt.Printf("Search Time (seconds): %f\n", diff.Seconds())
			break
		}
	}
}

func MatchFound(address, vanityKey string) bool {
	addressLowerCased := strings.ToLower(address)
	vanityKeyLowerCased := strings.ToLower(vanityKey)

	for i := 0; i < len(vanityKeyLowerCased); i++ {
		if addressLowerCased[i] != vanityKeyLowerCased[i] {
			return false
		}
	}

	return true
}
