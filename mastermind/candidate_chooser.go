package mastermind

import (
	"math/rand/v2"
)

// This is the strategy that governs how the candidate is chosen given a solution space
type candidateChooser interface {
	choose(solutionSpace []string) string
}

type RandomCandidateChooser struct{}

func (chooser RandomCandidateChooser) choose(solutionSpace []string) string {
	r := rand.IntN(len(solutionSpace))
	return solutionSpace[r]
}
