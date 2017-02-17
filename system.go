package voting

import (
	"demos/vox"
)

type System struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Info string `json:"info"`
	Vox string `json:"vox"`
	Criteria map[string]string `json:"criteria"`
}

const (
	UNKNOWN = "unknown"
	FALSE = "false"
	TRUE = "true"
	STRANGE = "strange"
	INAPPLICABLE = "false"
	STRANGE_YES = "strange-yes"
	STRANGE_NO = "strange-no"
)

var FIRST_PAST_THE_POST = System{
	Name : "First Past The Post",
	Url : "first-past-the-post",
	Info : FirstPastThePostInfo,
	Vox : vox.FirstPastThePost,
	Criteria: map[string]string{
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : FALSE,
		CONCORDET.Name : FALSE,
		CONCORDET_LOSER.Name : FALSE,
		SMITH_DOMINATED_ALTERNATIVES.Name : FALSE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : STRANGE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : TRUE,
		PARTICIPATION.Name : TRUE,
		REVERSAL_SYMMETRY.Name : FALSE,
		LATER_NO_HARM.Name : INAPPLICABLE,
		LATER_NO_HELP.Name : INAPPLICABLE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var INSTANT_RUN_OFF = System {
	Name : "Instant Runoff",
	Url : "instant-runoff",
	Info : InstantRunoffInfo,
	Vox : vox.InstantRunoff,
	Criteria: map[string]string {
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : TRUE,
		CONCORDET.Name : FALSE,
		CONCORDET_LOSER.Name : TRUE,
		SMITH_DOMINATED_ALTERNATIVES.Name : FALSE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : TRUE,
		MONOTONICITY.Name: FALSE,
		CONSISTENCY.Name : FALSE,
		PARTICIPATION.Name : FALSE,
		REVERSAL_SYMMETRY.Name : FALSE,
		LATER_NO_HARM.Name : TRUE,
		LATER_NO_HELP.Name : TRUE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var APPROVAL_VOTING = System {
	Name : "Approval Voting",
	Url : "approval-voting",
	Info: ApprovalVotingInfo,
	Vox : vox.Approval,
	Criteria: map[string]string{
		MAJORITY.Name : STRANGE,
		MUTUAL_MAJORITY.Name : FALSE,
		CONCORDET.Name : STRANGE,
		CONCORDET_LOSER.Name : FALSE,
		SMITH_DOMINATED_ALTERNATIVES.Name : STRANGE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : STRANGE_YES,
		IRRELEVANT_ALTERNATIVES.Name : STRANGE_YES,
		CLONE_ALTERNATIVES.Name : STRANGE_YES,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : STRANGE_YES,
		PARTICIPATION.Name : STRANGE_YES,
		REVERSAL_SYMMETRY.Name : TRUE,
		LATER_NO_HARM.Name : FALSE,
		LATER_NO_HELP.Name : STRANGE_YES,
		FAVOURITE_BETRAYAL.Name : TRUE,
	},
}

var SCORE_VOTING = System {
	Name : "Score Voting",
	Url : "score-voting",
	Info: ScoreVotingInfo,
	Vox : vox.Score,
	Criteria: map[string]string{

	},
}

var BORDA_COUNT = System {
	Name : "Borda Count",
	Url : "borda-count",
	Info : BordaCountInfo,
	Vox : vox.BordaCount,
	Criteria: map[string]string{
		MAJORITY.Name : FALSE,
		MUTUAL_MAJORITY.Name : FALSE,
		CONCORDET.Name : FALSE,
		CONCORDET_LOSER.Name : TRUE,
		SMITH_DOMINATED_ALTERNATIVES.Name : FALSE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : STRANGE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : TRUE,
		PARTICIPATION.Name : TRUE,
		REVERSAL_SYMMETRY.Name : TRUE,
		LATER_NO_HARM.Name : FALSE,
		LATER_NO_HELP.Name : TRUE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var SCHULZE_METHOD = System {
	Name : "Schulze Method",
	Url : "schulze-method",
	Info : SchulzeInfo,
	Vox : vox.Schulze,
	Criteria: map[string]string{
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : TRUE,
		CONCORDET.Name : TRUE,
		CONCORDET_LOSER.Name : TRUE,
		SMITH_DOMINATED_ALTERNATIVES.Name : TRUE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : TRUE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : FALSE,
		PARTICIPATION.Name : STRANGE,
		REVERSAL_SYMMETRY.Name : TRUE,
		LATER_NO_HARM.Name : FALSE,
		LATER_NO_HELP.Name : FALSE,
		FAVOURITE_BETRAYAL.Name : STRANGE,
	},
}

var KEMENY_YOUNG = System {
	Name : "Kemeny-Young",
	Url : "kemeny-young",
	Info : KemenyYoungInfo,
	Vox : vox.KemenyYoung,
	Criteria: map[string]string{
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : TRUE,
		CONCORDET.Name : TRUE,
		CONCORDET_LOSER.Name : TRUE,
		SMITH_DOMINATED_ALTERNATIVES.Name : TRUE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : TRUE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : STRANGE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : STRANGE,
		PARTICIPATION.Name : FALSE,
		REVERSAL_SYMMETRY.Name : TRUE,
		LATER_NO_HARM.Name : FALSE,
		LATER_NO_HELP.Name : FALSE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var MINIMAX = System {
	Name : "Minimax",
	Url : "minimax",
	Info : MinimaxInfo,
	Vox : vox.Minimax,
	Criteria: map[string]string{
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : FALSE,
		CONCORDET.Name : STRANGE,
		CONCORDET_LOSER.Name : FALSE,
		SMITH_DOMINATED_ALTERNATIVES.Name : FALSE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : STRANGE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : FALSE,
		PARTICIPATION.Name : FALSE,
		REVERSAL_SYMMETRY.Name : FALSE,
		LATER_NO_HARM.Name : STRANGE,
		LATER_NO_HELP.Name : FALSE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var COPELAND = System {
	Name : "Copeland",
	Url : "copeland",
	Info : CopelandInfo,
	Vox : vox.Copeland,
	Criteria: map[string]string{
		MAJORITY.Name : TRUE,
		MUTUAL_MAJORITY.Name : TRUE,
		CONCORDET.Name : TRUE,
		CONCORDET_LOSER.Name : TRUE,
		SMITH_DOMINATED_ALTERNATIVES.Name : TRUE,
		LOCAL_IRRELEVANT_ALTERNATIVES.Name : FALSE,
		IRRELEVANT_ALTERNATIVES.Name : FALSE,
		CLONE_ALTERNATIVES.Name : STRANGE,
		MONOTONICITY.Name: TRUE,
		CONSISTENCY.Name : FALSE,
		PARTICIPATION.Name : FALSE,
		REVERSAL_SYMMETRY.Name : TRUE,
		LATER_NO_HARM.Name : FALSE,
		LATER_NO_HELP.Name : FALSE,
		FAVOURITE_BETRAYAL.Name : FALSE,
	},
}

var (
	AllVotingSystems = []System{FIRST_PAST_THE_POST, INSTANT_RUN_OFF, APPROVAL_VOTING,
		SCORE_VOTING, BORDA_COUNT, SCHULZE_METHOD, COPELAND, MINIMAX, KEMENY_YOUNG}
)