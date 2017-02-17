package voting

type Vote struct {
	Rank map[string]int
	Signature []byte
}

type Ballot struct {

}

// VInput handles the parameters passed in by the user (i.e. not necessary)
// Victory method must answer a simple question --> should we end the election?
// VProgress tracks each stage of the voting process

// VOutput handles the graceful return o

type VotingMethod func(VotingInput) VotingOutput
type QuotaMethod func(int, int) int
type TieBreakMethod func(int, []VotingProgress, []string) []string

type VoteContainer interface {
	Input() VotingInput
	Voting() VotingMethod
	Output() VotingOutput
}

type VotingInput struct {
	NumElected int
	Candidates []string
	Quota QuotaMethod
	Votes func () []Vote
	TieBreak TieBreakMethod
	Distribute func ()
}

type VotingOutput struct {
	Notifications []string
	Stages []VotingProgress
	Results []string
}

func CountVotes(vc VoteContainer) (VotingOutput) {
	return vc.Voting()(vc.Input())
}

type VotingStage map[string][]Vote
type VotingProgress map[string]int




