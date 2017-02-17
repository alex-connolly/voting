package voting

func maximum(vp VotingProgress) []string {
	vals := make([]string, 0)
	max := -1
	for k, v := range vp {
		if v > max || v == -1 {
			max = v
			vals = []string{k}
		} else if v == max {
			vals = append(vals, k)
		}
	}
	return vals
}

func maxFloat(data map[string]float64) []string {
	vals := make([]string, 0)
	max := 0.0
	for k, v := range data {
		if v > max || v == 0.0 {
			max = v
			vals = []string{k}
		} else if v == max {
			vals = append(vals, k)
		}
	}
	return vals
}


func minimum(vp VotingProgress) [] string {
	vals := make([]string, 0)
	min := -1 // test
	for k, v := range vp {
		if v < min || min == -1 {
			min = v
			vals = []string{k}
		} else if v == min {
			vals = append(vals, k)
		}
	}
	return vals
}
