package test

import (
	"testing"
	"demos/voting"
)

func TestCoombsMethod(t *testing.T) {
	var container TestContainer
	container.vote = voting.CoombsMethod

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "nashville"){
		t.Error("wrong result")
	}
	if len(output.Stages) != 4 {
		t.Error(output.Stages)
	}
}

