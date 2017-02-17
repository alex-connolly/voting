package voting

import "demos/vox"

type Tie struct {
	Name string `json:"name"`
	Info string `json:"info"`
	Vox string `json:"vox"`
}

var RANDOM = Tie {
	Name : "Random Choice",
	Info : "Chooses a random selection to break ties.",
	Vox : vox.Random,
}

var AllTies = []Tie{RANDOM}
