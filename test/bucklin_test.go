package test

import (
	"testing"
	"demos/voting"
	"log"
)

func TestBucklin(t *testing.T){
	var container TestContainer
	container.vote = voting.Bucklin

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "memphis"){ // TODO: untested
		t.Error(output.Results[0])
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
	log.Println(output.Stages)
}
