package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/pepa65/mm/mastermind"
)

const (
	version = "0.4.0"
	pegs    = 8
	colors  = 10
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("mm v%s - Mastermind\nUsage:  mm [secret]\n", version)
		os.Exit(1)
	}

	Colors := "0123456789ABCDEFGKLMNOPQRSTUVXXYZ"[:colors]
	var secret string
	if len(os.Args) == 1 { // Generate a secret
		n := pegs
		for {
			secret += string(Colors[rand.IntN(colors)])
			n--
			if n < 1 {
				break
			}
		}
	} else { // Use given secret
		secret = os.Args[1]
	}
	fmt.Printf("Secret:   %s\n", secret)
	game := mastermind.Game{
		Pegs:    pegs,
		Colors:  Colors,
		Secret:  secret,
		Chooser: &mastermind.RandomCandidateChooser{},
	}
	if err := game.Solve(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	os.Exit(0)
}
