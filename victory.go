package voting

/*func MostVotes(max int, stage VotingStage) VotingResult {
	return nil
}*/

/*func Highlander(sum int, v VotingProgress) VotingResult {
	return VotingProgress{len(v) == 1, []string{"hi"}}
}

func Majority(sum int, v VotingProgress) VotingResult {
	return greaterThanProportion(0.5, sum, v)
}

/*func Quota(sum int, v VotingProgress) VotingResult {

}

func SuperMajority(sum int, v VotingProgress) VotingResult {
	// need to pass a percentage here
}

func greaterThanProportion(proportion float64, sum int, points VotingProgress) VotingResult {
	values := make([]string, 0)
	for k, v := range points {
		if float64(v) / float64(sum) > proportion {
			values = append(values, k)
		}
	}
	return VotingResult{len(values) > 0, values}
}

func Plurality(sum int, v VotingProgress) VotingResult {
	vals := make([]string, 0)
	max := -1
	for k, v := range v {
		if v > max {
			max = v
			vals = []string{k}
		} else if v == max {
			vals = append(vals, k)
		}
	}
	return VotingResult{true, vals} // doesn't handle two with the same
}*/



