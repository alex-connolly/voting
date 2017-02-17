package test

import (
	"testing"
	"demos/voting"
)

func TestSchulze(t *testing.T){
	var container TestContainer
	container.vote = voting.SchulzeMethod

	var input voting.VotingInput
	input.NumElected = 1
	input.Candidates = []string{"A", "B", "C", "D", "E"}
	input.Votes = getSchulzeData
	input.TieBreak = voting.ChooseRandom
	container.in = input

	output := voting.CountVotes(container)

	if len(output.Results) != 1 {
		t.Error(output.Results)
	}
	if !(output.Results[0] == "E"){
		t.Error(output.Results[0])
	}
}

func getSchulzeData() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"A": 1, "C": 2, "B":3, "E":4, "D":5}, nil}
	v2 := voting.Vote{map[string]int{"A": 1, "D": 2, "E":3, "C":4, "B":5}, nil}
	v3 := voting.Vote{map[string]int{"B": 1, "E": 2, "D":3, "A":4, "C":5}, nil}
	v4 := voting.Vote{map[string]int{"C": 1, "A": 2, "B":3, "E":4, "D":5}, nil}
	v5 := voting.Vote{map[string]int{"C": 1, "A": 2, "E":3, "B":4, "D":5}, nil}
	v6 := voting.Vote{map[string]int{"C": 1, "B": 2, "A":3, "D":4, "E":5}, nil}
	v7 := voting.Vote{map[string]int{"D": 1, "C": 2, "E":3, "B":4, "A":5}, nil}
	v8 := voting.Vote{map[string]int{"E": 1, "B": 2, "A":3, "D":4, "C":5}, nil}
	votes = append(votes, copyVote(5, v1)...)
	votes = append(votes, copyVote(5, v2)...)
	votes = append(votes, copyVote(8, v3)...)
	votes = append(votes, copyVote(3, v4)...)
	votes = append(votes, copyVote(7, v5)...)
	votes = append(votes, copyVote(2, v6)...)
	votes = append(votes, copyVote(7, v7)...)
	votes = append(votes, copyVote(8, v8)...)
	return votes
}