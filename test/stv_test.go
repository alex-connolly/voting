package test

import (
	"testing"
	"demos/voting"
)

func TestRandomChoiceSTV(t *testing.T){
	var container TestContainer
	container.vote = voting.RandomChoiceSTV

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}

func TestHareClarkSTV(t *testing.T){
	var container TestContainer
	container.vote = voting.HareClarkSTV

	var input voting.VotingInput
	input.NumElected = 3
	input.Votes = makeFoodElection
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 3 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 3 {
		t.Error(output.Stages)
	}
}

func TestGregorySTV(t *testing.T){
	var container TestContainer
	container.vote = voting.GregorySTV

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 1 {
		t.Error(output.Stages)
	}
}

func TestCPOSTVOne(t *testing.T){
	var container TestContainer
	container.vote = voting.CPOSTV

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"E", "A", "N", "G"}
	input.Votes = makeFoodElection
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 3 {
		t.Error(output.Stages)
	}
}

func TestCPOSTVTwo(t *testing.T){
	var container TestContainer
	container.vote = voting.CPOSTV

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"A", "B", "C", "D"}
	input.Votes = makeFoodElection
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 3 {
		t.Error(output.Stages)
	}
}

func TestConnollySTV(t *testing.T){
	var container TestContainer
	container.vote = voting.ConnollySTV

	var input voting.VotingInput
	input.NumElected = 3
	input.Votes = makeFoodElection
	input.TieBreak = voting.ChooseRandom
	input.Quota = voting.DroopQuota
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 3 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 3 {
		t.Error(output.Stages)
	}
}

func makeFoodElection() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"O": 1}, nil}
	v2 := voting.Vote{map[string]int{"P":1, "O":2}, nil}
	v3 := voting.Vote{map[string]int{"C": 1, "S":2}, nil}
	v4 := voting.Vote{map[string]int{"C": 1, "L":2}, nil}
	v5 := voting.Vote{map[string]int{"S": 1}, nil}
	v6 := voting.Vote{map[string]int{"L":1}, nil}
	votes = append(votes, copyVote(4, v1)...)
	votes = append(votes, copyVote(2, v2)...)
	votes = append(votes, copyVote(8, v3)...)
	votes = append(votes, copyVote(4, v4)...)
	votes = append(votes, copyVote(1, v5)...)
	votes = append(votes, copyVote(1, v6)...)
	return votes
}

func makeCPODatasetOne() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"E": 1}, nil}
	v2 := voting.Vote{map[string]int{"A":1, "N":2, "G":3}, nil}
	v3 := voting.Vote{map[string]int{"N": 1, "G":2}, nil}
	v4 := voting.Vote{map[string]int{"G": 1, "N":2}, nil}
	v5 := voting.Vote{map[string]int{"G": 1, "B":2}, nil}
	v6 := voting.Vote{map[string]int{"B":1, "G":2}, nil}
	votes = append(votes, copyVote(100, v1)...)
	votes = append(votes, copyVote(110, v2)...)
	votes = append(votes, copyVote(18, v3)...)
	votes = append(votes, copyVote(21, v4)...)
	votes = append(votes, copyVote(6, v5)...)
	votes = append(votes, copyVote(45, v6)...)
	return votes
}

func makeCPODatasetTwo() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"A": 1, "B":2, "C":3, "D": 4}, nil}
	v2 := voting.Vote{map[string]int{"A": 1, "C":2, "B":3, "D": 4}, nil}
	v3 := voting.Vote{map[string]int{"D":1}, nil}
	votes = append(votes, copyVote(5, v1)...)
	votes = append(votes, copyVote(17, v2)...)
	votes = append(votes, copyVote(8, v3)...)
	return votes
}