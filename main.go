package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pepa65/mm/mastermind"
)

const (
	version = "0.3.1"
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
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(os.Args) == 1 { // Generate a secret
		n := pegs
		for {
			secret += string(Colors[rnd.Intn(colors)])
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
