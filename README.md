# project-euler
Another implementation of Project Euler Problems :wink: #ProjectEuler100

## motivation

I just came across one freecodecamp [post](https://www.freecodecamp.org/news/projecteuler100-coding-challenge-competitive-programming/), which tempted me to accept _#ProjectEuler100_ challenge. It's not like that prior to this I never thought of solving these beautiful mathematical problems, but I never opensourced it. So it's looks like a great opportunity to me, where I can challenge my thinking capability.

## solutions

I'm planning to stick to _GoLang_ as language of implementation. All solutions will stay in this [directory](./projecteuler/).

### problem 1

#### statement

_If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23._

_Find the sum of all the multiples of 3 or 5 below 1000._

#### solution

233168

#### explanation

Iterating using a simple for loop _( starting from 3 )_, upto _(X -  1)_, where _X_ is given, and checking divisibility of current number by either 3 or 5. If it's divisible, then we add it up to _sum_ variable. And finally return _sum_, holding expected output.

**More coming soon ...** :wink:
