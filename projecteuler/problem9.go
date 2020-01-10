package projecteuler

// IsPythagoreanTriplet - Given three natural numbers, a, b, c
// we check whether a < b < c & a^2 + b^2 == c^2 or not
func IsPythagoreanTriplet(a int, b int, c int) bool {
	return a < b && b < c && (a*a+b*b) == c*c
}

// SpecialPythagoreanTriplet - A pythrorean triplet
// which satisfies a + b + c = 1000, there's exactly one of this kind
//
// we start iterating from c=999 ( outer loop ), b=998, a=997 ( inner most loop )
// satisfying a<b<c always. Now we need to check whether a+b+c == 1000, if it's lesser than so,
// breaking out of loop, avoiding unnecessary computations.
//
// If previous condition and pythagorean triplet condition is satisfied
// we calculate product of a, b, c, which is our desired result.
func SpecialPythagoreanTriplet() int {
	prod := 1
	for c := 999; c > 0; c-- {
		for b := c - 1; b > 0; b-- {
			for a := b - 1; a > 0; a-- {
				sum := a + b + c
				if sum > 1000 {
					continue
				} else if sum == 1000 && IsPythagoreanTriplet(a, b, c) {
					prod = a * b * c
					break
				} else {
					break
				}
			}
			if prod != 1 {
				break
			}
		}
		if prod != 1 {
			break
		}
	}
	return prod
}
