package test

import (
	"testing"
	"demos/voting"
)

func TestInstantRunOff(t *testing.T){
	var container TestContainer
	container.vote = voting.InstantRunoff

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "knoxville") {
		t.Error("wrong result")
	}
	if len(output.Results) != 1 {
		t.Error(len(output.Results))
	}
	if len(output.Stages) != 4 {
		t.Error(output.Stages)
	}
}

func TestInstantRunOffMultipleWinners(t *testing.T){
	var container TestContainer
	container.vote = voting.InstantRunoff

	var input voting.VotingInput
	input.NumElected = 2
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "knoxville") {
		t.Error("wrong result")
	}
	if len(output.Results) != 2 {
		t.Error(len(output.Results))
	}
	if len(output.Stages) != 4 {
		t.Error(len(output.Stages))
	}
}

func TestInstantRunoffNotEnoughCandidates(t *testing.T){
	var container TestContainer
	container.vote = voting.InstantRunoff

	var input voting.VotingInput
	input.NumElected = 5
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if !(output.Results[0] == "knoxville") {
		t.Error("wrong result")
	}
	if len(output.Results) != 4 { // 4 and not 5 (only 4 candidates)
		t.Error(len(output.Results))
	}
	if len(output.Stages) != 4 {
		t.Error(len(output.Stages))
	}
}

func TestInstantRunoffFullyTied(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeTiedWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
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

func TestInstantRunoffMultipleFullyTied(t *testing.T) {
	var container TestContainer
	container.vote = voting.FirstPastThePost

	var input voting.VotingInput
	input.NumElected = 2
	input.Votes = makeTiedWikipediaVotes
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

func TestInstantRunoffMultiplePartiallyTied(t *testing.T) {
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