package voting

import "demos/vox"

type Quota struct {
	Name string `json:"name"`
	Info string `json:"info"`
	Formula string `json:"formula"`
	Vox string `json:"vox"`
}

var IMPERIALI = Quota {
	Name : "Imperiali",
	Info : "This is a quota",
	Vox : vox.Imperiali,
}

var HAGENBACH_BISCHOFF = Quota {
	Name : "Hagenbach Bischoff",
	Info : "This is a quota",
	Vox : vox.HagenbachBischoff,
}

var HARE_QUOTA = Quota {
	Name : "Hare",
	Info : "This is also a quota",
	Vox : vox.Hare,
}

var DROOP_QUOTA = Quota {
	Name : "Droop",
	Info : "Another quota",
	Vox : vox.Droop,
}

var AllQuotas = []Quota{IMPERIALI, HAGENBACH_BISCHOFF, HARE_QUOTA, DROOP_QUOTA}