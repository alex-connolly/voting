package test

import (
	"testing"
	"demos/voting"
)

func TestAntiPlurality(t *testing.T){
	var container TestContainer
	container.vote = voting.AntiPlurality

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if !(output.Results[0] == "nashville") && !(output.Results[0] == "chattanooga"){
		t.Error(output.Results[0])
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}
