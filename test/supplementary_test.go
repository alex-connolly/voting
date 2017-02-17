package test

import (
	"testing"
	"demos/voting"
)

func TestSupplementary(t *testing.T){
	var container TestContainer
	container.vote = voting.SupplementaryVoting

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}
