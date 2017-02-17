package voting

import (
	"strconv"
	"math"
	"fmt"
	"sort"
	"log"
	"math/rand"
)

func createOutput() VotingOutput {
	var out VotingOutput
	out.Stages = make([]VotingProgress, 0)
	out.Notifications = make([]string, 0)
	return out
}

func FirstPastThePost(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	count := 0
	for count < input.NumElected {
		if len(data) == 0 {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - count) + " of the " + strconv.Itoa(count) + " slots.")
		}
		temp := convertStageToProgress(data)
		res := maximum(temp)
		// handle tiebreak instances
		if count + len(res) > input.NumElected {
			res = input.TieBreak(input.NumElected - count, out.Stages, res)
		}
		count += len(res)
		out.Results = append(out.Results, res...)
		out.Stages = append(out.Stages, temp)
		// remove all of the victors from the data
		for _, s := range res {
			out.Notifications = append(out.Notifications, s + " elected with " + strconv.Itoa(temp[s]) + " votes.")
			delete(data, s)
		}
	}
	return out
}

func InstantRunoff(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	order := make([]tempResult, 0)
	for len(data) != 0 {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		res := minimum(temp)
		 // it is bad to lose the tie break (you are eliminated)
		// 3 elected, 4 left, 2 tied, if
		if len(res) > 1 && len(data) - len(res) < input.NumElected { //TODO: CHECK LOGIC (could be v wrong)
			res = input.TieBreak(input.NumElected - len(data), out.Stages, res)
		}
		distribute(data, res)
		for _, o := range res {
			order = append(order, tempResult{o, temp[o]})
		}
	}
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

type tempResult struct {
	name string
	number int
}

// currently requires that all ballots are fully filled-in
func BordaCount(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	points := borda(votes)
	count := 0
	for count < input.NumElected {
		out.Stages = append(out.Stages, points)
		res := maximum(points)
		if count + len(res) > input.NumElected {
			res = input.TieBreak(input.NumElected - count, out.Stages, res)
		}
		for _, o := range res {
			delete(points, o)
		}
		out.Results = append(out.Results, res...)
		count += len(res)
	}
	return out
}

func MajorityJudgement(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := make(map[string]map[int]int)
	for _, vote := range votes {
		for k, v := range vote.Rank {
			if data[k] == nil {
				data[k] = make(map[int]int)
			}
			data[k][v]++
		}
	}
	// currently doesn't handle ties
	min, median := getMediansAndMins(data)
	for len(min) > 1 && len(data) > 0 {
		for _, m := range min {
			data[m][median]--
		}
		min, median = getMediansAndMins(data)
	}
	out.Results = min
	return out
}

func getMediansAndMins(data map[string]map[int]int) ([]string, int) {
	medians := make(map[string]int)
	for k, v := range data {
		mk := make([]int, len(v))
		i := 0
		for k, _ := range v {
			mk[i] = k
			i++
		}
		sort.Sort(sort.Reverse(sort.IntSlice(mk)))
		length := 0
		for _, number := range mk {
			length += v[number]
		}

		progress := 0
		for _, number := range mk {
			if float64(progress + v[number]) >= 0.5 * float64(length) {
				medians[k] = number
				break
			}
			progress += v[number]
		}
	}
	minMedian := 0
	min := []string{}
	for k, v := range medians {
		if minMedian > v || minMedian == 0 {
			minMedian = v
			min = []string{k}
		} else if minMedian == v {
			min = append(min, k)
		}
	}
	return min, minMedian
}

func ApprovalVoting(input VotingInput) VotingOutput {
	out := createOutput()
	scoreSums := make(map[string]int)
	votes := input.Votes()
	for _, vote := range votes {
		for opt, _ := range vote.Rank {
			scoreSums[opt]++
		}
	}
	count := 0
	for count < input.NumElected {
		if len(scoreSums) == 0 {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - count) + " of the " + strconv.Itoa(count) + " slots.")
		}
		res := maximum(scoreSums)
		// handle tiebreak instances
		if count + len(res) > input.NumElected {
			res = input.TieBreak(input.NumElected - count, out.Stages, res)
		}
		count += len(res)
		out.Results = append(out.Results, res...)
		// remove all of the victors from the data
		for _, s := range res {
			out.Notifications = append(out.Notifications, s + " elected with " + strconv.Itoa(scoreSums[s]) + " votes.")
			delete(scoreSums, s)
		}
	}
	out.Stages = []VotingProgress{scoreSums}
	return out
}

func ScoreVoting(input VotingInput) VotingOutput {
	out := createOutput()
	scoreSums := make(map[string]int)
	votes := input.Votes()
	for _, vote := range votes {
		for opt, val := range vote.Rank {
			scoreSums[opt] += val
		}
	}
	count := 0
	for count < input.NumElected {
		if len(scoreSums) == 0 {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - count) + " of the " + strconv.Itoa(count) + " slots.")
		}
		res := maximum(scoreSums)
		// handle tiebreak instances
		if count + len(res) > input.NumElected {
			res = input.TieBreak(input.NumElected - count, out.Stages, res)
		}
		count += len(res)
		out.Results = append(out.Results, res...)
		out.Stages = append(out.Stages, scoreSums)
		// remove all of the victors from the data
		for _, s := range res {
			out.Notifications = append(out.Notifications, s + " elected with " + strconv.Itoa(scoreSums[s]) + " votes.")
			delete(scoreSums, s)
		}
	}
	out.Stages = []VotingProgress{scoreSums}
	return out
}

func borda(votes []Vote) VotingProgress {
	points := make(map[string]int)
	for _, vote := range votes {
		for opt, val := range vote.Rank {
			points[opt] += len(vote.Rank) - val
		}
	}
	return points
}

func SchulzeMethod(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	matrix := make([][]float64, len(input.Candidates)) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range matrix {
		matrix[i] = make([]float64, len(input.Candidates))
	}

	for _, vote := range votes {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix); j++ {
				if vote.Rank[input.Candidates[i]] > vote.Rank[input.Candidates[j]] {
					matrix[i][j]++
				}
			}
		}

	}
	paths := make([][]float64, len(input.Candidates)) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range paths {
		paths[i] = make([]float64, len(input.Candidates))
	}
	for h := 0; h < len(input.Candidates); h++ {
		for w := 0; w < len(input.Candidates); w++ {
			for i := 0; i < len(input.Candidates); i++ {
				for j := 0; j < len(input.Candidates); j++ {
					if i != j {
						if matrix[i][j] > matrix[j][i] {
							paths[i][j] = matrix[i][j]
						} else {
							paths[i][j] = 0
						}
					}
				}
			}
			for i := 0; i < len(input.Candidates); i++ {
				for j := 0; j < len(input.Candidates); j++ {
					if i != j {
						for k := 0; k < len(input.Candidates); k++ {
							if i != k && j != k {
								paths[j][k] = math.Max(paths[j][k], math.Min(paths[j][i], paths[i][k]))
							}
						}
					}
				}
			}
		}
	}
	// not sure if this is the optimal way to do this
	better := make(map[string]int)
	for i := 0; i < len(input.Candidates); i++ {
		for j := 0; j < len(input.Candidates); j++ {
			if paths[j][i] > paths[i][j] {
				better[input.Candidates[i]]++
			}
		}
	}
	order := SortedKeys(better)
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(out.Results)])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func pos(vals []string, value string) int {
	for p, v := range vals {
		if (v == value) {
			return p
		}
	}
	return -1
}

/*func BordaNanson(input VotingInput, out VotingOutput) VotingOutput {
	votes := input.Votes()
	points := borda(votes)
	out.Stages = []VotingProgress{points}
	order := make([]tempResult, 0)
	if len(points) != 0 {
		excluded := belowOrEqualToAverage(points)
		delete(points, excluded...)
		temp := borda(votes)
		out.Stages = append(out.Stages, temp)
	}
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func BordaBaldwin(input VotingInput, out VotingOutput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	dist := countFirstPreferences(votes)
	points := borda(votes)
	out.Stages = []VotingProgress{points}
	out.Result = victory(len(votes), out.Stages[0])
	if !out.Result.Success {
		dist = exclude(dist, minimum(dist))
		temp := borda(votes)
		for k, v := range temp {
			points[k] += v
		}
		out.Stages = append(out.Stages, points)
		out.Result = victory(len(votes), convertStageToProgress(dist))
	}
	return out
}*/


func SriLankanContingentVote(input VotingInput) VotingOutput {
	return SupplementaryVoting(input)
}


// terrible for actual elections (super vulnerable to tactical voting)
func AntiPlurality(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	lasts := make(map[string]int)
	// only works on fully filled in ballots
	for _, vote := range votes {
		for k, v := range vote.Rank {
			if v == len(input.Candidates) {
				lasts[k]++
			}
		}
	}
	// make sure our minimum includes those candidates who were not placed last at all
	for _, c := range input.Candidates {
		if lasts[c] == 0 { // if it has the default 0 value
			lasts[c] = 0 // actually set it to 0
		}
	}
	out.Stages = []VotingProgress{lasts}
	log.Println(lasts)
	for len(out.Results) < input.NumElected {
		if len(lasts) > len(out.Results){
			min := minimum(lasts)
			if len(min) + len(out.Results) > input.NumElected {
				min = input.TieBreak(input.NumElected - len(out.Results), out.Stages, min)
			}
			out.Results = append(out.Results, min...)
			for _, m := range min {
				delete(lasts, m)
			}
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

/*func RankedPairs(votes []Vote, victory func (int, VotingProgress) VotingResult) ([]VotingProgress, VotingResult){

}

func ExhaustiveBallot(votes []Vote, victory func (int, VotingProgress) VotingResult) ([]VotingProgress, VotingResult){

}

func RankedPairs(votes []Vote, victory func (int, VotingProgress) VotingResult) ([]VotingProgress, VotingResult){

}*/

func CPOSTV(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	quota := input.Quota(len(votes), input.NumElected)
	winners := make([]string, 0)
	distributing := true
	for distributing {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		// should be in order of descending surplus
		vals := SortedKeys(temp)
		distributing = false
		for _, v := range vals{
			if temp[v] >= quota {
				distributing = true
				surplus := temp[v] - quota
				surplusVotes := make(map[string]int)
				for _, vote := range data[v] {
					valid := findNearestValid(vote.Rank[v], vote, data)
					if valid != "" {
						surplusVotes[valid]++
					}
				}
				ratio := float64(surplus) / float64(len(surplusVotes))
				if ratio > 1 {
					ratio = 1
				}
				log.Println("ratio")
				log.Println(ratio)
				log.Println(surplus)
				log.Println(len(surplusVotes))
				for k, v := range surplusVotes {
					for i := 0; float64(i) < float64(v) * ratio; i++ {
						data[k] = append(data[k], Vote{map[string]int{k:1}, nil})
					}
				}
			}
		}
	}
	waiting := make([]string, 0)
	for i := 0; i < len(input.Candidates); i++ {
		if !sliceContains(winners, input.Candidates[i]) {
			waiting = append(waiting, input.Candidates[i])
		}
	}

	perms := permutations(waiting)
	matrix := make([][]int, len(perms))
	for i := range matrix {
		matrix[i] = make([]int, len(perms))
	}
	permMatrix := make([][][]string, len(perms))
	for i := range permMatrix {
		permMatrix[i] = make([][]string, len(perms))
		for j := range permMatrix[i] {
			permMatrix[i][j] = perms[i]
		}
	}
	// 1. Eliminate all candidates who are not in either outcome.
	// almost definitely a way to optimise this
	for i := 0; i < len(permMatrix); i++ {
		for j := 0; j < len(permMatrix); j++ {
			dataCopy := make(map[string][]Vote)
			for k, v := range data {
				dataCopy[k] = v
			}
			for k, v := range dataCopy {
				if !sliceContains(permMatrix[i][j], k) && !sliceContains(permMatrix[j][i], k) {
					delete(dataCopy, k)
				}
				if sliceContains(permMatrix[i][j], k) && sliceContains(permMatrix[j][i], k) {
					if len(v) > quota {
						//dataCopy := input.Distribute()
					}
				}
				var scoreOne, scoreTwo int
				for _, vals := range permMatrix[i][j]{
					scoreOne += len(dataCopy[vals])
				}
				for _, vals := range permMatrix[j][i]{
					scoreTwo += len(dataCopy[vals])
				}
				if scoreOne > scoreTwo {
					matrix[i][j]++
					matrix[j][i]--
				}
				if scoreTwo > scoreOne {
					matrix[j][i]++
					matrix[i][j]--
				}
			}
		}
	}
	log.Println(matrix)
	return out
}

/*func STVCLE(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	quota := input.Quota(len(votes), input.NumElected)
	condorcet()
	return out
}*/

func HareClarkSTV(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	winners := make([]tempResult, 0)
	quota := input.Quota(len(votes), input.NumElected)
	for len(winners) < input.NumElected {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		// should be in order of descending surplus
		vals := SortedKeys(temp)
		for _, v := range vals{
			if temp[v] >= quota {
				log.Println(v + " reached quota")
				surplus := temp[v] - quota
				surplusVotes := make(map[string]int)
				winners = append(winners, tempResult{v, temp[v]})
				for _, vote := range data[v] {
					valid := findNearestValid(vote.Rank[v], vote, data)
					if valid != "" {
						surplusVotes[valid]++
					}
				}
				delete(data, v)
				ratio := float64(surplus) / float64(len(surplusVotes))
				if ratio > 1 {
					ratio = 1
				}
				log.Println("ratio")
				log.Println(ratio)
				log.Println(surplus)
				log.Println(len(surplusVotes))
				for k, v := range surplusVotes {
					for i := 0; float64(i) < float64(v) * ratio; i++{
						data[k] = append(data[k], Vote{map[string]int{k:1}, nil})
					}
				}
			}
		}
		temp = convertStageToProgress(data)
		res := minimum(temp)
		// it is bad to lose the tie break (you are eliminated)
		// 3 elected, 4 left, 2 tied, if
		if len(res) > 1 && len(data) - len(res) < input.NumElected { //TODO: CHECK LOGIC (could be v wrong)
			res = input.TieBreak(input.NumElected - len(data), out.Stages, res)
		}
		distribute(data, res)

	}
	for len(out.Results) < input.NumElected {
		if len(winners) > len(out.Results){
			out.Results = append(out.Results, winners[len(out.Results)].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func GregorySTV(input VotingInput) VotingOutput {
	out := createOutput()
	return out
}

func ConnollySTV(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	winners := make([]tempResult, 0)
	quota := input.Quota(len(votes), input.NumElected)
	for len(data) != 0 {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		for k, v := range temp {
			if v >= quota {
				winners = append(winners, tempResult{k, quota})
				quotaVotes := make(map[string][]Vote)
				for _, vote := range data[k] {
					val := findNearestValid(vote.Rank[k], vote, data)
					if quotaVotes[val] == nil {
						quotaVotes[val] = make([]Vote, 0)
					}
					quotaVotes[val] = append(quotaVotes[val], vote)
				}
				keys := make([]string, 0)
				for k, _ := range quotaVotes {
					keys = append(keys, k)
				}
				matrix := concordet(keys, votes)
				// there will be a better way to do this i.e. not O(n!) brute force
				ranking := make(map[int][]string)
				copied := make([]string, len(keys))
				for i, v := range keys {
					copied[i] = v
				}
				for _, perm := range permutations(copied) {
					count := 0
					for i, val := range perm {
						count += countLower(val, matrix, perm[i:], keys)
					}
					ranking[count] = perm
				}
				log.Println(ranking)
				// distribute according to ranking
				sum := 0
				for k, _ := range ranking {
					sum += k
				}
				log.Println(sum)
				for rank, order := range ranking {
					for i := 0; i < rank / sum; i++ {
						var v Vote
						v.Rank = make(map[string]int)
						count := 1
						for _, val := range order {
							if data[val] != nil{
								v.Rank[val] = count
								count++
							}
						}
						data[order[0]] = append(data[order[0]], v)

					}
				}

			}
		}
		temp = convertStageToProgress(data)
		res := minimum(temp)
		// it is bad to lose the tie break (you are eliminated)
		// 3 elected, 4 left, 2 tied, if
		if len(res) > 1 && len(data) - len(res) < input.NumElected { //TODO: CHECK LOGIC (could be v wrong)
			res = input.TieBreak(input.NumElected - len(data), out.Stages, res)
		}
		distribute(data, res)


	}

	for len(out.Results) < input.NumElected {
		if len(winners) > len(out.Results){
			out.Results = append(out.Results, winners[len(out.Results)].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func RandomChoiceSTV(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	order := make([]tempResult, 0)
	quota := input.Quota(len(votes), input.NumElected)
	for len(order) < input.NumElected {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		// should be in order of descending surplus
		vals := SortedKeys(temp)
		log.Println(temp)
		for _, v := range vals {
			if temp[v] >= quota {
				surplusSize := temp[v] - quota
				random := getRandomVotes(surplusSize, votes)
				for _, vote := range random {
					val := findNearestValid(vote.Rank[v], vote, data)
					data[val] = append(data[val], vote) // check corner cases
				}
			}
		}
		res := minimum(temp)
		// it is bad to lose the tie break (you are eliminated)
		// 3 elected, 4 left, 2 tied, if
		if len(res) > 1 && len(data) - len(res) < input.NumElected { //TODO: CHECK LOGIC (could be v wrong)
			res = input.TieBreak(input.NumElected - len(data), out.Stages, res)
		}
		distribute(data, res)
		for _, o := range res {
			order = append(order, tempResult{o, temp[o]})
		}

	}
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

// returns a slice of continuous votes (wrapping around) from the original slice
func getRandomVotes(number int, votes []Vote) []Vote {
	random := make([]Vote, 0)
	bottom := rand.Intn(len(votes))
	// e.g. need 2, bottom 1, length 4
	if bottom + number < len(votes){
		for i := bottom; i < bottom + number; i++ {
			random = append(random, votes[i])
		}
	} else {
		// e.g. need 2, bottom 3, length 4
		for i := bottom; i < len(votes); i++ {
			random = append(random, votes[i])
		}
		for i := 0; len(random) < number; i++ {
			random= append(random, votes[i])
		}
	}
	return votes
}

func BTRSTV(){

}

func STVCLE(){

}

func hareClarkRatio(totalVotes, quota, exhausted int) float64 {
	return float64(quota - totalVotes) / float64(totalVotes - exhausted)
}

func AlternativeVote(input VotingInput) VotingOutput {
	out := createOutput()
	return out
}

func Bucklin(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	points := convertStageToProgress(data)
	count := 1
	for !candidateHasMajority(len(votes), points) {
		for _, v := range data {
			for _, vote := range v {
				for opt, val := range vote.Rank {
					if val == count {
						points[opt]++
					}
				}
			}
		}
		count++
		out.Stages = append(out.Stages, points)
	}
	// TODO: do we want to exclude candidates with 0 first prefs if majority found on first iteration
	// what if this candidate should get a seat
	for len(out.Results) < input.NumElected {
		if len(points) > 0 {
			max := maximum(points)
			if len(out.Results) + len(max) > input.NumElected {
				max = input.TieBreak(input.NumElected - len(out.Results), out.Stages, max)
			}
			out.Results = append(out.Results, max...)
			for _, m := range max {
				delete(points, m)
			}
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - count) + " of the " + strconv.Itoa(count) + " slots.")
			break
		}
	}
	return out
}

//TODO: check if there is a need to check for a majority after each round
func CoombsMethod(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	order := make([]tempResult, 0)
	for len(data) != 0 {
		temp := convertStageToProgress(data)
		out.Stages = append(out.Stages, temp)
		res := mostLastPlaces(data)
		// it is bad to lose the tie break (you are eliminated)
		// 3 elected, 4 left, 2 tied, if
		if len(res) > 1 && len(data) - len(res) < input.NumElected { //TODO: CHECK LOGIC (could be v wrong)
			res = input.TieBreak(input.NumElected - len(data), out.Stages, res)
		}
		distribute(data, res)
		for _, o := range res {
			order = append(order, tempResult{o, temp[o]})
		}
	}
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1].name)
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func candidateHasMajority(sum int, p VotingProgress) bool {
	for _, v := range p {
		if float64(v) / float64(sum) > 0.5 {
			return true
		}
	}
	return false
}

func belowOrEqualToAverage(p VotingProgress) [] string {
	avg := average(p)
	excluded := make([]string, 0)
	for k, v := range p {
		if float64(v) <= avg {
			excluded = append(excluded, k)
		}
	}
	return excluded
}


func mostLastPlaces(data map[string][]Vote) [] string {
	lasts := make(map[string]int)
	for _, votes := range data {
		for _, vote := range votes {
			min := -1
			minEl := ""
			for opt, val := range vote.Rank {
				if _, ok := data[opt]; ok {
					if len(vote.Rank) - val < min || min == -1 {
						min = len(vote.Rank) - val
						minEl = opt
					}
				}
			}
			lasts[minEl]++
			// TODO: ensure that this is how it works
			// i.e. counting the lowest valid (not necessarily last if last is already gone)
		}
	}
	return maximum(lasts)
}



func SupplementaryVoting(input VotingInput) VotingOutput {
	out := createOutput()
	votes := input.Votes()
	data := countFirstPreferences(votes)
	points := convertStageToProgress(data)
	out.Stages = []VotingProgress{points}
	var first, second string
	max := maximum(points)
	// will break if 0 results
	if len(max) > 2{
		max = input.TieBreak(2, out.Stages, max)
		first, second = max[0], max[1]
	} else if len(max) == 1{
		first = max[0]
		delete(points, max[0])
		max = maximum(points)
		if len(max) > 1 {
			second = input.TieBreak(1, out.Stages, max)[0]
		}
	} else {
		first, second = max[0], max[1]
	}
	// just redo all votes, will go to the right places anyway
	var firstScore, secondScore int
	for _, vote := range votes {
		firstInt, secondInt := 1000000, 1000000
		for k, v := range vote.Rank {
			if k == first {
				firstInt = v
			} else if k == second {
				secondInt = v
			}
		}
		if firstInt > secondInt {
			firstScore++
		} else if firstInt < secondInt {
			secondScore++
		}
	}
	if firstScore > secondScore {
		out.Results = append(out.Results, first)
	} else if secondScore > firstScore {
		out.Results = append(out.Results, second)
	} else {
		out.Results = append(out.Results, input.TieBreak(1, out.Stages, []string{first, second})...)
	}
	return out

}

func concordet(candidates []string, votes []Vote) [][]int {
	matrix := make([][]int, len(candidates))
	for i := range matrix {
		matrix[i] = make([]int, len(candidates))
	}
	for _, vote := range votes {
		for i := 0; i < len(candidates); i++ {
			for j := 0; j < len(candidates); j++ {
				// what happens if not found??
				if vote.Rank[candidates[i]] == 0 && vote.Rank[candidates[j]] == 0{

				} else if vote.Rank[candidates[i]] == 0{

				} else if vote.Rank[candidates[j]] == 0 {
					matrix[j][i]++
				} else {
					if vote.Rank[candidates[i]] < vote.Rank[candidates[j]] {
						matrix[j][i]++
					}
				}
			}
		}
	}
	return matrix
}

func Condorcet(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	matrix := concordet(input.Candidates, votes)
	better := make(map[string]int)
	for i := 0; i < len(input.Candidates); i++ {
		for j := 0; j < len(input.Candidates); j++ {
			if matrix[j][i] > matrix[i][j] {
				better[input.Candidates[i]]++
			}
		}
	}
	order := SortedKeys(better)
	if better[order[0]] == len(input.Candidates) - 1 {
		out.Results = append(out.Results, order[0])
	}
	return out
}

func CondorcetKemenyYoung(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	matrix := concordet(input.Candidates, votes)
	// there will be a better way to do this i.e. not O(n!) brute force
	ranking := make(map[int][]string)
	copied := make([]string, len(input.Candidates))
	for i, v := range input.Candidates {
		copied[i] = v
	}
	for _, perm := range permutations(copied) {
		count := 0
		for i, val := range perm {
			count += countLower(val, matrix, perm[i:], input.Candidates)
		}
		ranking[count] = perm
	}

	max := 0
	var maxEl []string
	for k, v := range ranking {
		if k > max {
			max = k
			maxEl = v
		}
	}
	for len(out.Results) < input.NumElected {
		if len(maxEl) > len(out.Results){
			out.Results = append(out.Results, maxEl[len(out.Results)])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func countLower(pref string, matrix [][]int, than, candidates []string) int {
	sum := 0
	for _, v := range than {
		//log.Println(pref + " with " + v)
		//log.Println("index: [" + strconv.Itoa(pos(candidates, v)) + "," +  strconv.Itoa(pos(candidates, pref)) + "]")
		sum += matrix[pos(candidates, v)][pos(candidates, pref)]
		//log.Println("sum: " + strconv.Itoa(sum))
	}
	return sum
}


func first(data sort.Interface) {
	sort.Sort(data)
}

// next returns false when it cannot permute any more
// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func next(data sort.Interface) bool {
	var k, l int
	for k = data.Len() - 2; ; k-- {
		if k < 0 {
			return false
		}
		if data.Less(k, k+1) {
			break
		}
	}
	for l = data.Len() - 1; !data.Less(k, l); l-- {
	}
	data.Swap(k, l)
	for i, j := k+1, data.Len()-1; i < j; i++ {
		data.Swap(i, j)
		j--
	}
	return true
}

// permuteStrings returns all possible permutations of string slice.
func permutations(slice []string) [][]string {
	first(sort.StringSlice(slice))

	copied1 := make([]string, len(slice)) // we need to make a copy!
	copy(copied1, slice)
	result := [][]string{copied1}

	for {
		isDone := next(sort.StringSlice(slice))
		if !isDone {
			break
		}

		// https://groups.google.com/d/msg/golang-nuts/ApXxTALc4vk/z1-2g1AH9jQJ
		// Lesson from Dave Cheney:
		// A slice is just a pointer to the underlying back array, your storing multiple
		// copies of the slice header, but they all point to the same backing array.

		// NOT
		// result = append(result, slice)

		copied2 := make([]string, len(slice))
		copy(copied2, slice)
		result = append(result, copied2)
	}

	combNum := 1
	for i := 0; i < len(slice); i++ {
		combNum *= i + 1
	}
	if len(result) != combNum {
		fmt.Printf("Expected %d combinations but %+v because of duplicate elements", combNum, result)
	}

	return result
}

func CondorcetMiniMaxWorstPairwiseOpposition(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	worstPairwiseOpposition := make(map[string]int)
	matrix := concordet(input.Candidates, votes)
	// uses worst pairwise opposition
	for i := 0; i < len(input.Candidates); i++ {
		max := 0
		for j := 0; j < len(input.Candidates); j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
		worstPairwiseOpposition[input.Candidates[i]] = max
	}
	order := SortedKeys(worstPairwiseOpposition)
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func CondorcetMiniMaxWorstMargins(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	margins := make(map[string]int)
	matrix := concordet(input.Candidates, votes)
	// uses worst pairwise opposition
	for i := 0; i < len(input.Candidates); i++ {
		max := 0
		for j := 0; j < len(input.Candidates); j++ {
			if matrix[i][j] - matrix[j][i] > max {
				max = matrix[i][j] - matrix[j][i]
			}
		}
		margins[input.Candidates[i]] = max
	}
	order := SortedKeys(margins)
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func CondorcetMiniMaxWorstWinningVotes(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	winning := make(map[string]int)
	matrix := concordet(input.Candidates, votes)
	// uses worst pairwise opposition
	for i := 0; i < len(input.Candidates); i++ {
		max := 0
		for j := 0; j < len(input.Candidates); j++ {
			if matrix [i][j] > matrix[j][i] {
				if matrix[i][j] > max {
					max = matrix[i][j]
				}
			}
		}
		winning[input.Candidates[i]] = max
	}
	order := SortedKeys(winning)
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(order)-len(out.Results)-1])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}

func CondorcetCopeland(input VotingInput) VotingOutput{
	out := createOutput()
	votes := input.Votes()
	matrix := concordet(input.Candidates, votes)
	wins := make(map[string]int)
	// loop structure ensures each match up occurs only once
	for i := 0; i < len(input.Candidates); i++ {
		for j := i + 1; j < len(input.Candidates); j++ {
			if matrix [i][j] < matrix[j][i] {
				wins[input.Candidates[i]]++
				wins[input.Candidates[j]]--
			} else if matrix [i][j] > matrix[j][i] {
				wins[input.Candidates[i]]--
				wins[input.Candidates[j]]++
			}
		}
	}
	order := SortedKeys(wins)
	// TODO: add tiebreaking
	for len(out.Results) < input.NumElected {
		if len(order) > len(out.Results){
			out.Results = append(out.Results, order[len(out.Results)])
		} else {
			out.Notifications = append(out.Notifications, "Not enough candidates to fill " + strconv.Itoa(input.NumElected - len(out.Results)) + " of the " + strconv.Itoa(input.NumElected) + " slots.")
			break
		}
	}
	return out
}


