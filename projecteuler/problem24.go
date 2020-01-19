package projecteuler

import "strings"

// swaps two string elements using their address ( i.e. pointer )
func swap(a *string, b *string) {
	tmp := *a
	*a = *b
	*b = tmp
}

// reverses a given string sequence, from starting index to end index
func reverse(arr []string, start int, end int) {
	for i := 0; i < (end-start+1)/2; i++ {
		swap(&arr[start+i], &arr[end-i])
	}
}

// Generates lexicographic next term ( string ), using Narayana's Algorithm, 14th Century
func generateNextPerm(arr []string) {
	last := len(arr) - 1 // last index of set
	if last < 1 {        // if we've only one element in set, nothing to permute
		return
	}
	i := last - 1
	for i >= 0 && !(arr[i] < arr[i+1]) { // finds out highest index ( i ) in set such that arr[i] < arr[i+1]
		i--
	}
	if i < 0 { // if it's lesser than 0, we've generated last possible sequence, reverting back to initial combination
		reverse(arr, 0, last)
	} else {
		j := last
		for j > i+1 && !(arr[j] > arr[i]) { // finding out index ( j ) in set such that arr[i] < arr[j]
			j--
		}
		swap(&arr[i], &arr[j])  // swapping arr[i] & arr[j]
		reverse(arr, i+1, last) // reversing all elements of set, starting from index `i+1` to `last`
	}
	// and now we've next lexicographic permutation / reverted back to initial combination
}

// LexicographicPermutations - Returns 10^6-th lexicographic permuation
// of a given slice of strings (0-9)
func LexicographicPermutations() string {
	arr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	c := 1
	for c < 1000000 {
		generateNextPerm(arr)
		c++
	}
	return strings.Join(arr, "")
}
