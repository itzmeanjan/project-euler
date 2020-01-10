# project-euler
Another implementation of Project Euler Problems :wink: #ProjectEuler100

## motivation

I just came across one freecodecamp [post](https://www.freecodecamp.org/news/projecteuler100-coding-challenge-competitive-programming/), which tempted me to accept _#ProjectEuler100_ challenge. It's not like that prior to this I never thought of solving these beautiful mathematical problems, but I never opensourced it. So it's looks like a great opportunity to me, where I can challenge my thinking capability.

## solutions

I'm planning to stick to _GoLang_ as language of implementation. All solutions will stay in this [directory](./projecteuler/).

### [problem 1](./projecteuler/problem1.go)

#### statement

_If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23._

_Find the sum of all the multiples of 3 or 5 below 1000._

#### solution

233168 in 7.325µs

#### explanation

Iterating using a simple for loop _( starting from 3 )_, upto _(X -  1)_, where _X_ is given, and checking divisibility of current number by either 3 or 5. If it's divisible, then we add it up to _sum_ variable. And finally return _sum_, holding expected output.

### [problem 2](./projecteuler/problem2.go)

#### statement

_Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:_

_1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ..._

_By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms._

#### solution

4613732 in 2.752µs

#### explanation

Using dynamic programming style, for calculating fibonacci terms, recursive strategy will be straightforward but run slow ( and no doubt very expensive ). Starting with a slice of two elements `{1, 2}`, we'll keep calculating next fibonacci term until most recently computed term crosses _4,000,000_. And in each iteration, it'll check whether this term is even or not. If even, we'll add it up to _sum_ variable, which is initialized with _2_ ( because at very beginning _fibArr_, was only holding _2_ as even number )

### [problem 3](./projecteuler/problem3.go)

#### statement

_The prime factors of 13195 are 5, 7, 13 and 29._

_What is the largest prime factor of the number 600851475143 ?_

#### solution

6857 in 171.188338ms

#### explanation

First calculates square root of given number, and find out all primes which are under or equals to that sqrt value. Now we'll simply iterate over that prime holder slice, from last to first, i.e. from higher value prime to lower value prime, cause finally, we need to find out maximum prime factor of _num_. That'll allow us to perform lesser number of checkings.

Generation of primes under _X_, is done using dynamic programming strategy, by updating a slice holding primes, on runtime. Because we know any composite number must have prime factor, lesser than square root of that number. So we'll perform check with prime numbers only, which will save a lot of computation too.

### [problem 4](./projecteuler/problem4.go)

#### statement

_A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99._

_Find the largest palindrome made from the product of two 3-digit numbers._

#### solution

906609 in 172.4µs

#### explanation

We'll start from end i.e. for finding largest possible palindrome number under _1000_, we'll start checking from _999_ & keep multiplying two numbers _( < 1000 )_, until I reach _1_. But that'll be brute-force, which is why we'd prefer breaking out of current iteration, as soon as current product _( product in this iteration )_ goes below `largestPalim` _( which is largest palindrome computed upto this point )_.

### [problem 5](./projecteuler/problem5.go)

#### statement

_2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder._

_What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?_

#### solution

232792560 in 1.664408611s

#### explanation

We'll start finding smallest number divisible by all numbers from _1_ to _20_, at 10, because for any number to be divisible by _10_, it must end with round figure. And keep incrementing number under lens by _10_ ( after each iteration ), which will eventually reduce #-of computational steps required for finding result.

### [problem 6](./projecteuler/problem6.go)

#### statement

_The sum of the squares of the first ten natural numbers is, 12 + 22 + ... + 102 = 385_

_The square of the sum of the first ten natural numbers is,(1 + 2 + ... + 10)2 = 552 = 3025_

_Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640._

_Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum._

#### solution

25164150 in 275ns

#### explanation

We'll calculate square of sum of {1..100} ( used _n*(n+1)/2_, for finding sum of first _n_ natural numbers ) & square of each natural number {1..100}, while accumulating them up in a single variable. Finally a simple absolute substraction of those two, will get us desired result.

### [problem 7](./projecteuler/problem7.go)

#### statement

_By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13._

_What is the 10,001st prime number?_

#### solution

104743 in 14.523821ms

#### explanation

We'll buffer all primes calculated uptil now, and check a certain odd number's _( for reducing number of steps, we're skipping even numbers, cause they are definitely composite )_ divisibility using primes _( buffered )_ under square root of `num`. Finally we return _(x-1)_ indexed term from buffer.

### [problem 8](./projecteuler/problem8.go)

#### statement

_The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832._

_73167176531330624919225119674426574742355349194934_
_96983520312774506326239578318016984801869478851843_
_85861560789112949495459501737958331952853208805511_
_12540698747158523863050715693290963295227443043557_
_66896648950445244523161731856403098711121722383113_
_62229893423380308135336276614282806444486645238749_
_30358907296290491560440772390713810515859307960866_
_70172427121883998797908792274921901699720888093776_
_65727333001053367881220235421809751254540594752243_
_52584907711670556013604839586446706324415722155397_
_53697817977846174064955149290862569321978468622482_
_83972241375657056057490261407972968652414535100474_
_82166370484403199890008895243450658541227588666881_
_16427171479924442928230863465674813919123162824586_
_17866458359124566529476545682848912883142607690042_
_24219022671055626321111109370544217506941658960408_
_07198403850962455444362981230987879927244284909188_
_84580156166097919133875499200524063689912560717606_
_05886116467109405077541002256983155200055935729725_
_71636269561882670428252483600823257530420752963450_

_Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?_

#### solution

23514624000 in 36.693µs

#### explanation

Given a _1000_ digit number _( as a string )_, we'll iterate over all indices of this string _( from 0 to 999 )_, so that we can consider all possible consequtive substrings of fixed length. Then we keep multiplying all digits present in each sub-sequence and take max of it. That'll satisfy our need, but multiplication of digits in subsequences is overlapping problem. So, we won't recompute, but use previous iterations cached value _( with care )_, if and only if it's not initial iteration & first digit of previous subsequence isn't _0_.

**More coming soon ...** :wink:
