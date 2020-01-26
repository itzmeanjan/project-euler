package projecteuler

// using dynamic programming, we'll cache precalculated values so that
// recomputation can be avoided
func count(coins [8]int, total int) int {
	cache := make([]int, total+1)     // taking a caching table for holding precomputed values
	cache[0] = 1                      // for making 0 pence, we've only 1 way, this is base condition
	for i := 0; i < len(coins); i++ { // in each iteration picking a single coin
		for j := coins[i]; j <= total; j++ { // and try to figure out whether this can be used for making target amount i.e. `total`
			cache[j] += cache[j-coins[i]]
		}
	}
	return cache[total] // returning final amount, place at `total`-th index,
	// holding possible number of ways for making `total` pence ( target amount )
}

// CoinSum - How many different ways 2 Pound = 200 pence can be made,
// using only given coins
func CoinSum() int {
	return count([8]int{1, 2, 5, 10, 20, 50, 100, 200}, 200)
}
