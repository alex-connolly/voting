package voting

import "math"

func DroopQuota(sum, winners int) int {
	return int(math.Floor((float64(sum) / float64(winners + 1)) + 1))
}

func HagenbachBischoffQuota(sum, winners int) int {
	return int(float64(sum) / float64(winners) + 1)
}

func HareQuota(sum, winners int) int {
	return int(float64(sum)/float64(winners))
}

func ImperialiQuota(sum, winners int) int {
	return int(float64(sum) / float64(winners) + 2)
}