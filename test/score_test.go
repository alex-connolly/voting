package test

import (
	"testing"
	"demos/voting"
)

func TestScore(t *testing.T){
	var container TestContainer
	container.vote = voting.ScoreVoting

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

func TestApproval(t *testing.T){
	var container TestContainer
	container.vote = voting.ApprovalVoting

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

func TestMajorityJudgement(t *testing.T){
	var container TestContainer
	container.vote = voting.MajorityJudgement

	var input voting.VotingInput
	input.NumElected = 1
	input.Votes = makeWikipediaVotes
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)
	if output.Results[0] != "nashville" {
		t.Error(output.Results)
	}
	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if len(output.Stages) != 0 {
		t.Error(output.Stages)
	}
}
