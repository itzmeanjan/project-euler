package main

import (
	"fmt"
	"time"

	"github.com/itzmeanjan/project-euler/projecteuler"
)

func main() {
	fmt.Println("\n:: #ProjectEuler100 ::")
	start := time.Now()
	fmt.Printf("\n\tProblem 1 : %d in %v\n", projecteuler.FindSumOfAllMultiplesOf3or5BelowX(1000), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 2 : %d in %v\n", projecteuler.GetSumOfEvenValuedFibonacciTermsUnderX(4000000), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 3 : %d in %v\n", projecteuler.GetLargestPrimeFactor(600851475143), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 4 : %d in %v\n", projecteuler.GetLargestPalindrome(1000), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 5 : %d in %v\n", projecteuler.GetSmallestMultiple(), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 6 : %d in %v\n", projecteuler.SumSquareDiff(100), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 7 : %d in %v\n", projecteuler.GetXthPrime(10001), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 8 : %d in %v\n", projecteuler.LargestProductInSeries(13), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("\n\tProblem 9 : %d in %v\n", projecteuler.SpecialPythagoreanTriplet(), time.Now().Sub(start))
	fmt.Println("Done")
}
