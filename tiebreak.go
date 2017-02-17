package voting

import (
	"math/rand"
)

// chooses 1 of the result values at random
func ChooseRandom(number int, stages []VotingProgress, vr []string) []string {
	bottom := rand.Intn(len(vr))
	vals := make([]string, 0)
	// e.g. need 2, bottom 1, length 4
	if bottom + number < len(vr){
		for i := bottom; i < bottom + number; i++ {
			vals = append(vals, vr[i])
		}
	} else {
		// e.g. need 2, bottom 3, length 4
		for i := bottom; i < len(vr); i++ {
			vals = append(vals, vr[i])
		}
		for i := 0; len(vals) < number; i++ {
			vals = append(vals, vr[i])
		}
	}
	return vals
}



/*func excludeBottomProportion(proportion float64, data VotingStage) VotingStage {
	//excludedOptions := make([]string, 0)
}*/

/*func excludeBelowProportion(proportion float64, data VotingStage) VotingStage {
	sum := sum(data)
	excludedOptions := make([]string, 0)
	for k, v := range data {
		if float64(len(v)) / float64(sum) < proportion {
			excludedOptions = append(excludedOptions, k)
		}
	}
	for _, opt := range excludedOptions {
		exclude(data, opt)
	}
	return data
}*/




