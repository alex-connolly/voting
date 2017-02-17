package test

import (
	"testing"
	"demos/voting"
)

func TestBordaCount(t *testing.T){
	var container TestContainer
	container.vote = voting.BordaCount

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "nashville"){
		t.Error("wrong result")
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}
/*func TestBordaNanson(t *testing.T){
	var container TestContainer
	container.vote = voting.BordaNanson

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
	if len(output.Stages) != 2 {
		t.Error(output.Stages)
	}
}

func TestBordaBaldwin(t *testing.T){
	var container TestContainer
	container.vote = voting.BordaBaldwin

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "nashville"){
		t.Error("wrong result")
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}*/
