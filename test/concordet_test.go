package test

import (
	"testing"
	"demos/voting"
)

func TestConcordet(t *testing.T){
	var container TestContainer
	container.vote = voting.Condorcet

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "knoxville", "chattanooga"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetKemenyYoung(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetKemenyYoung

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetMiniMaxWorstPairwiseOpposition(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetMiniMaxWorstPairwiseOpposition

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetMiniMaxWorstMargin(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetMiniMaxWorstMargins

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetMiniMaxWorstWinningVotes(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetMiniMaxWorstWinningVotes

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetDodgson(t *testing.T){

}

func TestConcordetCopeland(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetCopeland

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"memphis", "nashville", "chattanooga", "knoxville"}
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "nashville"){
		t.Error(output.Results[0])
	}
}

func TestConcordetCopelandAlternate(t *testing.T){
	var container TestContainer
	container.vote = voting.CondorcetCopeland

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"A", "B", "C", "D", "E"}
	input.Votes = makeAlternateWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "A"){
		t.Error(output.Results[0])
	}
}