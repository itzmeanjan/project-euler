package main

import (
	"fmt"

	"github.com/itzmeanjan/project-euler/projecteuler"
)

func main() {
	fmt.Println("\n:: #ProjectEuler100 ::")
	fmt.Printf("\n\tProblem 1 : %d\n", projecteuler.FindSumOfAllMultiplesOf3or5BelowX(1000))
	fmt.Printf("\n\tProblem 2 : %d\n", projecteuler.GetSumOfEvenValuedFibonacciTermsUnderX(4000000))
	fmt.Printf("\n\tProblem 3 : %d\n", projecteuler.GetLargestPrimeFactor(600851475143))
}
