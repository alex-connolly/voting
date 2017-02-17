package voting

const (
	MonotonicityInfo = "A ranked voting system is monotonic if it is neither possible to prevent " +
	"the election of a candidate by ranking them higher on some of the ballots, nor possible to " +
	"elect an otherwise unelected candidate by ranking them lower on some of the ballots " +
	"(while nothing else is altered on any ballot)."
	ConcordetInfo = "A system satisties the Condorcet criterion if a candidate who would win in a " +
		"direct contest with all other candidates will always win the election."
	ConcordetLoserInfo = "A system satisties the Condorcet criterion if a candidate who would lose in a " +
		"direct contest with all other candidates will always lose the election."
	ReversalSymmetryInfo = ""
	LaterNoHarmInfo = ""
	LaterNoHelpInfo = ""
	BetrayalInfo = ""
	ParticipationInfo = ""
	ConsistencyInfo = ""
	CloneproofInfo = ""
	LIIAInfo = ""
	IIAInfo = ""
	ISDAInfo = ""
	MutualMajorityInfo = ""
	MajorityInfo = ""
)

const (
	FirstPastThePostInfo = "FPTP system."
	InstantRunoffInfo = "IRV voting."
	ApprovalVotingInfo = "Approval voting."
	ScoreVotingInfo = "Score voting."
	BordaCountInfo = ""
	KemenyYoungInfo = ""
	CopelandInfo = ""
	MinimaxInfo = ""
	SchulzeInfo = ""
)
