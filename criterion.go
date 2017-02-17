package voting

type Criterion struct {
	Name string `json:"name"`
	Shortened string `json:"shortened"`
	Url string `json:"url"`
	Info string `json:"info"`
}

var MONOTONICITY = Criterion {
	Name : "Monotonicity",
	Url : "monotonicity",
	Info : MonotonicityInfo,
}

var CONCORDET = Criterion {
	Name : "Concordet",
	Url : "concordet",
	Info: ConcordetInfo,
}

var CONCORDET_LOSER = Criterion {
	Name : "Concordet Loser",
	Url : "concordet-loser",
	Info: ConcordetLoserInfo,
}

var MAJORITY = Criterion {
	Name : "Majority",
	Url : "majority",
	Info : MajorityInfo,
}

var MUTUAL_MAJORITY = Criterion {
	Name : "Mutual Majority",
	Url : "mutual-majority",
	Info : MutualMajorityInfo,
}

var SMITH_DOMINATED_ALTERNATIVES = Criterion {
	Name : "Independence of Smith Dominated Alternatives",
	Shortened : "ISDA",
	Url : "isda",
	Info : ISDAInfo,
}

var IRRELEVANT_ALTERNATIVES = Criterion {
	Name : "Independence of Irrelevant Alternatives",
	Shortened : "IIA",
	Url : "iia",
	Info : IIAInfo,
}

var LOCAL_IRRELEVANT_ALTERNATIVES = Criterion {
	Name : "Local Independence of Irrelevant Alternatives",
	Shortened: "LIIA",
	Url: "liia",
	Info: LIIAInfo,
}

var CLONE_ALTERNATIVES = Criterion {
	Name : "Independence of Clone Alternatives",
	Shortened: "Cloneproof",
	Url: "cloneproof",
	Info: CloneproofInfo,
}

var CONSISTENCY = Criterion {
	Name : "Consistency",
	Url : "consistency",
	Info : ConsistencyInfo,
}

var PARTICIPATION = Criterion {
	Name : "Participation",
	Url : "participation",
	Info : ParticipationInfo,
}

var REVERSAL_SYMMETRY = Criterion {
	Name : "Reversal Symmetry",
	Shortened: "Reversal",
	Url : "reversal-symmetry",
	Info : ReversalSymmetryInfo,
}

var LATER_NO_HARM = Criterion {
	Name : "Later No Harm",
	Url : "later-no-harm",
	Info : LaterNoHarmInfo,
}

var LATER_NO_HELP = Criterion {
	Name : "Later No Help",
	Url : "later-no-help",
	Info : LaterNoHelpInfo,
}

var FAVOURITE_BETRAYAL = Criterion {
	Name : "Favourite Betrayal",
	Shortened: "Betrayal",
	Url : "favourite-betrayal",
	Info : BetrayalInfo,
}

var (
	AllCriteria = []Criterion{
		MONOTONICITY, CONCORDET, CONCORDET_LOSER,
		MAJORITY, MUTUAL_MAJORITY, SMITH_DOMINATED_ALTERNATIVES,
		IRRELEVANT_ALTERNATIVES, LOCAL_IRRELEVANT_ALTERNATIVES,
		CLONE_ALTERNATIVES, CONSISTENCY, PARTICIPATION,
		LATER_NO_HARM, LATER_NO_HELP, FAVOURITE_BETRAYAL,
	}
)

