package main

import (
	"os"

	"github.com/adriangracia/mastering-bitcoin/chapter4"
)

func main() {
	chapter4.MineVanityAddress(os.Args[1])
}
