package test

import (
	"testing"
	"demos/voting"
)

/*func BenchmarkFirstPastThePost(b *testing.B) {
	votes := getRandomVotes(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstPastThePost(votes)
	}
}

func getRandomVotes(number int) []Vote {
	votes := make([]Vote, number)
	for i := 0; i < number; i++{
		var v Vote
		v.rank = make(map[string]int)
		v.rank[strconv.Itoa(rand.Intn(100))] = 1
		votes[i] = v
	}
	return votes
}   */

type TestContainer struct {
	in voting.VotingInput
	vote voting.VotingMethod
	out voting.VotingOutput
}

func (t TestContainer) Voting() voting.VotingMethod    {
	return t.vote
}

func (t TestContainer) Input() voting.VotingInput {
       return t.in
}

func (t TestContainer) Output() voting.VotingOutput {
	               return t.out
}

func TestMakeWikipediaVotes(t *testing.T){
	if len(makeWikipediaVotes()) != 100 {
		t.Error("wrong num of votes")
	}
}

func makeWikipediaVotes() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"memphis": 1, "nashville": 2, "chattanooga":3, "knoxville":4}, nil}
	v2 := voting.Vote{map[string]int{"nashville": 1, "chattanooga": 2, "knoxville":3, "memphis":4}, nil}
	v3 := voting.Vote{map[string]int{"chattanooga": 1, "knoxville": 2, "nashville":3, "memphis":4}, nil}
	v4 := voting.Vote{map[string]int{"knoxville": 1, "chattanooga": 2, "nashville":3, "memphis":4}, nil}
	votes = append(votes, copyVote(42, v1)...)
	votes = append(votes, copyVote(26, v2)...)
	votes = append(votes, copyVote(15, v3)...)
	votes = append(votes, copyVote(17, v4)...)
	return votes
}

func makeAlternateWikipediaVotes() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"A": 1, "E": 2, "C":3, "D":4, "B":5}, nil}
	v2 := voting.Vote{map[string]int{"B": 1, "A": 2, "E":3}, nil}
	v3 := voting.Vote{map[string]int{"C": 1, "D": 2, "B":3}, nil}
	v4 := voting.Vote{map[string]int{"D": 1, "A": 2, "E":3}, nil}
	votes = append(votes, copyVote(31, v1)...)
	votes = append(votes, copyVote(30, v2)...)
	votes = append(votes, copyVote(29, v3)...)
	votes = append(votes, copyVote(10, v4)...)
	return votes
}

func makeTiedWikipediaVotes() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"memphis": 1, "nashville": 2, "chattanooga":3, "knoxville":4}, nil}
	v2 := voting.Vote{map[string]int{"nashville": 1, "chattanooga": 2, "knoxville":3, "memphis":4}, nil}
	v3 := voting.Vote{map[string]int{"chattanooga": 1, "knoxville": 2, "nashville":3, "memphis":4}, nil}
	v4 := voting.Vote{map[string]int{"knoxville": 1, "chattanooga": 2, "nashville":3, "memphis":4}, nil}
	votes = append(votes, copyVote(42, v1)...)
	votes = append(votes, copyVote(42, v2)...)
	votes = append(votes, copyVote(42, v3)...)
	votes = append(votes, copyVote(42, v4)...)
	return votes
}

func makePartiallyTiedWikipediaVotes() []voting.Vote {
	votes := make([]voting.Vote, 0)
	v1 := voting.Vote{map[string]int{"memphis": 1, "nashville": 2, "chattanooga":3, "knoxville":4}, nil}
	v2 := voting.Vote{map[string]int{"nashville": 1, "chattanooga": 2, "knoxville":3, "memphis":4}, nil}
	v3 := voting.Vote{map[string]int{"chattanooga": 1, "knoxville": 2, "nashville":3, "memphis":4}, nil}
	v4 := voting.Vote{map[string]int{"knoxville": 1, "chattanooga": 2, "nashville":3, "memphis":4}, nil}
	votes = append(votes, copyVote(42, v1)...)
	votes = append(votes, copyVote(42, v2)...)
	votes = append(votes, copyVote(26, v3)...)
	votes = append(votes, copyVote(17, v4)...)
	return votes
}

func copyVote(times int, v voting.Vote) []voting.Vote {
	votes := make([]voting.Vote, times)
	for i := 0; i < times; i++ {
		votes[i] = v
	}
	return votes
}
