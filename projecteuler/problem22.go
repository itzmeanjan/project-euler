package projecteuler

import (
	"io/ioutil"
	"strings"
)

// calculates name score of a given string
// simply convert each character into its position number in
// alphabet, and sum it up
func namescore(name string) int {
	score := 0
	for _, i := range name {
		score += (int(i) - 65 + 1)
	}
	return score
}

// sorts a given slice of strings in ascendic manner
func sortNames(names []string) []string {
	for i := 0; i < len(names); i++ {
		for j := 0; j < i; j++ {
			if names[i] < names[j] {
				swap(&names[i], &names[j])
			}
		}
	}
	return names
}

// reads content of data file, holding 5K names
// and returns a slice of strings ( each element is a name )
func readFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tmp := strings.Split(string(content), ",")
	for i, j := range tmp {
		tmp[i] = strings.Trim(j, "\"")
	}
	return tmp
}

// NamesScores - Calculates total name score of a given
// data set ( holding 5K names )
func NamesScores() int {
	total := 0
	for i, j := range sortNames(readFile("p022_names.txt")) {
		total += (i + 1) * namescore(j)
	}
	return total
}
