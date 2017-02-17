package voting

import "sort"

func sumPoints(points map[string]int) int {
	sum := 0
	for _, v := range points {
		sum += v
	}
	return sum
}

func average(points map[string]int) float64 {
	return float64(sumPoints(points)) / float64(len(points))
}

func countFirstPreferences(votes []Vote) map[string][]Vote {
	vm := make(map[string][]Vote)
	for _, vote := range votes {
		for k, v := range vote.Rank {
			if v == 1 {
				if (vm[k] == nil){
					vm[k] = make([]Vote, 0)
				}
				vm[k] = append(vm[k], vote)
			}
		}
	}
	return vm
}

type SortedMap struct {
	M map[string]int
	S []string
}

func (sm *SortedMap) Len() int {
	return len(sm.M)
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.M[sm.S[i]] > sm.M[sm.S[j]]
}

func (sm *SortedMap) Swap(i, j int) {
	sm.S[i], sm.S[j] = sm.S[j], sm.S[i]
}

func SortedKeys(m map[string]int) []string {
	sm := new(SortedMap)
	sm.M = m
	sm.S = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.S[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.S
}

func distribute(data VotingStage, exclude []string) VotingStage {
	excludedVotes := make(map[string][]Vote)
	for _, s := range exclude {
		excludedVotes[s] = data[s]
		delete(data, s)
	}
	for key, votes := range excludedVotes {
		for _, vote := range votes {
			num := vote.Rank[key]
			k := findNearestValid(num, vote, data)
			if k != "" {
				data[k] = append(data[k], vote)
			}
		}
	}
	return data
}

func convertStageToProgress(stage VotingStage) VotingProgress {
	progress := make(map[string]int)
	for k, v := range stage {
		progress[k] = len(v)
	}
	return progress
}

func findNearestValid(rank int, vote Vote, stage VotingStage) string {
	min := 100000
	minEl := ""
	for k, v := range vote.Rank {
		if _, ok := stage[k]; ok {
			if v > rank && v - rank < min {
				min = v - rank
				if min == 1 {
					return k
				}
				minEl = k
			}
		}
	}
	return minEl
}

func sliceContains(data []string, point string) bool{
	for _, val := range data {
		if val == point {
			return true
		}
	}
	return false
}
