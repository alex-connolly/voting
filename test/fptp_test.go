package test

import (
	"testing"
	"demos/voting"
	"log"
)

func TestFirstPastThePost(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "memphis"){
		t.Error("wrong result")
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
	if len(output.Notifications) != 1 {
		t.Error(output.Notifications)
	}
}

func TestFirstPastThePostMultiple(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 3
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	log.Println(output.Results)
	if !(output.Results[0] == "memphis"){
		t.Error(output.Results[0])
	}
	if len(output.Stages) != 3 {
		t.Error(output.Stages)
	}
	if len(output.Notifications) != 3 {
		t.Error(output.Notifications)
	}
}

func TestFirstPastThePostFullyTied(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeTiedWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	log.Println(output.Results)
	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
	if len(output.Notifications) != 1 {
		t.Error(output.Notifications)
	}
}

func TestFirstPastThePostMultipleFullyTied(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 2
	input.Votes = makeTiedWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	log.Println(output.Results)
	if len(output.Results) != 2 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
	if len(output.Notifications) != 2 {
		t.Error(output.Notifications)
	}
}

func TestFirstPastThePostMultiplePartiallyTied(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 2
	input.Votes = makePartiallyTiedWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 2 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
	if len(output.Notifications) != 2 {
		t.Error(output.Notifications)
	}
}

