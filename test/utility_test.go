package test

import (
	"testing"
	"fmt"
	"demos/voting"
)

func TestSortedMap(t *testing.T){
	s := []string{"Python", "Python", "Python", "igor", "igor", "igor", "igor", "go", "go", "Golang", "Golang", "Golang", "Golang", "Py", "Py"}
	count := make(map[string]int)

	for _, v := range s {
		count[v]++
	}

	for _, res := range voting.SortedKeys(count) {
		fmt.Println(res, count[res])
	}
}
