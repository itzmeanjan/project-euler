package projecteuler

import (
	"strings"
)

type permHelper struct {
	elems       []string
	k           int
	orderedSeqs []string
}

func swap(a *string, b *string) {
	tmp := *a
	*a = *b
	*b = tmp
}

func (pr *permHelper) permute() {
	if len(pr.orderedSeqs) == 1000000 {
		return
	}
	if pr.k == 1 {
		pr.orderedSeqs = sortNames(append(pr.orderedSeqs, strings.Join(pr.elems, "")))
	} else {
		pr.k--
		pr.permute()
		pr.k++
		for i := 0; i < pr.k-1; i++ {
			if pr.k%2 == 0 {
				swap(&pr.elems[i], &pr.elems[pr.k-1])
			} else {
				swap(&pr.elems[0], &pr.elems[pr.k-1])
			}
			pr.k--
			pr.permute()
			pr.k++
		}
	}
}

// LexicographicPermutations - ...
func LexicographicPermutations() string {
	perm := permHelper{[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, 10, make([]string, 0)}
	perm.permute()
	return perm.orderedSeqs[999999]
}
