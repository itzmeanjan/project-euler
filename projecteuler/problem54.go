package projecteuler

import (
	"io/ioutil"
	"strings"
)

// holds what's in hand of a player
// presents five cards, which is to be sorted as per
// defined criteria
type hand struct {
	cards []string
}

// given a hand of five cards, it'll get us
// a mapping from card value to card suits
// if we've multiple cards of same value which
// belong to different suits, then we'll keep them
// in a container, for that certain card value
//
// e.g. we've a hand of cards [5C 5D 9C AD AC]
// which can be represented as a map {3:[C D] 7:[C] 12:[D C]}
// where 3, 7, 12 i.e. keys are numeric representation of card values, see more about it below
//
// now I've incorporated way to map only a subset of cards in hand,
// where subset is no doubt denoted by its starting and ending indices ( inclusive )
func (h hand) getMap(from int, to int) map[int][]string {
	buffer := make(map[int][]string)
	for i := from; i <= to; i++ {
		v := h.cards[i]
		tmp := cardValueToIndex(v[:1])
		if _, ok := buffer[tmp]; ok {
			buffer[tmp] = append(buffer[tmp], v[1:])
		} else {
			buffer[tmp] = []string{v[1:]}
		}
	}
	return buffer
}

// given a hand of cards, it'll sort them in ascending fashion
// as per card values i.e. 2 < 3 < 4 < 5 < 6 < 7 < 8 < 9 < 10 < Jack < Queen < King < Ace
func (h *hand) sort() {
	for i := 0; i < len(h.cards); i++ {
		for j := 0; j < len(h.cards)-i-1; j++ {
			if needsSwap(h.cards[j][:1], h.cards[j+1][:1]) {
				swap(&h.cards[j], &h.cards[j+1])
			}
		}
	}
}

/*
	From here on we'll start writing logic for
	determining whether a given hand ( combination
	of cards sorted in ascending fashion ) satisfies a property or not,
	using which we can determine, winner of current hand

	Well following 10 functions will help us in ranking
	a hand, which are presented in their ascending order
	of importance
*/

// finds highest card value for a subset of cards
// in hand now
//
// card values to be expressed in their corresponding numeric values
// starting from 1 to 13
func (h hand) getHighCard(from int, to int) int {
	max := 0
	for k := range h.getMap(from, to) {
		if k > max {
			max = k
		}
	}
	return max
}

// checks whether we've a pair of cards of same value
// in case of success returns card value, with which we've a pair of cards
// else returns 0, to denote abscence
func (h hand) getOnePair() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 2 && k > max {
			max = k
		}
	}
	return max
}

// checks whether we've two different pairs
// where each with same card value
// in case of success returns a slice of those card values
// in case of abscence returns `nil`
func (h hand) getTwoPairs() []int {
	pairs := []int{}
	for k, v := range h.getMap(0, 4) {
		if len(v) == 2 {
			pairs = append(pairs, k)
		}
	}
	if len(pairs) == 2 {
		return pairs
	}
	return nil
}

// finds card value for which we've three cards
// else returns 0, to denote abscence
func (h hand) getThreeOfAKind() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 3 && k > max {
			max = k
		}
	}
	return max
}

// checks whether all cards in hand
// are of consequtive value or not
func (h hand) isStraight() bool {
	check := true
	cur := cardValueToIndex(h.cards[0][:1])
	for i := 1; i < len(h.cards); i++ {
		if tmp := cardValueToIndex(h.cards[i][:1]); tmp != cur+1 {
			check = false
			break
		} else {
			cur = tmp
		}
	}
	return check
}

// checks whether all cards in hand are of same suit
// or not i.e. need to check second value of each card in hand
// e.g. for 5H, 9C, 6H , we'll check {H, C, H}
func (h hand) isFlush() bool {
	check := true
	cur := ""
	for _, v := range h.getMap(0, 4) {
		if len(v) != 1 {
			check = false
			break
		}
		if cur == "" {
			cur = v[0]
		} else if cur != v[0] {
			check = false
			break
		}
	}
	return check
}

// checks whether we've three cards of same kind
// and a pair of cards with same value, returns corresponding
// values respectively
func (h hand) getFullHouse() (int, int) {
	return h.getThreeOfAKind(), h.getOnePair()
}

// finds card value, with which we've four cards
// else returns 0
func (h hand) getFourOfAKind() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 4 && k > max {
			max = k
		}
	}
	return max
}

// checks whether all cards are of consecutive values
// and of same suit
//
// whether they belong to same suit can be checked using
// isFlush() function, defined above
func (h hand) isStraightFlush() bool {
	return h.isStraight() && h.isFlush()
}

// chceks whether we've 10, Jack, Queen, King, Ace
// and all are of same suit
//
// whether they belong to same suit can be checked using
// isFlush() function, defined above
func (h hand) isRoyalFlush() bool {
	count := 0
	for _, v := range h.cards {
		if tmp := cardValueToIndex(v[1:]); !(tmp >= 9 && tmp <= 13) {
			break
		}
		count++
	}
	return count == 5 && h.isFlush()
}

// given a hand of cards, we'll compute its rank
// depending upon above defined 10 functions
// so it's pretty evident we'll have 10 ranks
// starting from 1 and running upto 10
//
// we'll it may be the situation, when hand
// satisfies rank 3 & 2 criteria both,
// then we'll choose higher rank, cause these
// ranks will help us in deciding winner
// between two players
func (h hand) getRank() int {
	if h.isRoyalFlush() {
		return 10
	}
	if h.isStraightFlush() {
		return 9
	}
	if h.getFourOfAKind() != 0 {
		return 8
	}
	if three, pair := h.getFullHouse(); three != 0 && pair != 0 {
		return 7
	}
	if h.isFlush() {
		return 6
	}
	if h.isStraight() {
		return 5
	}
	if h.getThreeOfAKind() != 0 {
		return 4
	}
	if h.getTwoPairs() != nil {
		return 3
	}
	if h.getOnePair() != 0 {
		return 2
	}
	return 1
}

// given a card value, returns corresponding numeric
// representation, which will be really helpful in comparing
// two cards or while sorting cards ascendically
// i.e. maps {2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace}
// to {1, 2, ..., 12, 13}
func cardValueToIndex(a string) int {
	cardVToIdxMap := map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
	return cardVToIdxMap[a]
}

// given two card values, determines whether we need
// to swap them in order to obtain a ascendically
// sorted card sequence or not
//
// corresponding numeric values of cards to be used for comparison
func needsSwap(a string, b string) bool {
	if aIdx, bIdx := cardValueToIndex(a), cardValueToIndex(b); aIdx < bIdx {
		return false
	} else if aIdx == bIdx {
		return false
	}
	return true
}

// splits hand to two players into two different hands
// where first one is for player one & second one for
// player two.
// i.e. processes each line of input data file
func getHandsOfTwoPlayers(hands string) (hand, hand) {
	tmp := strings.Split(hands, " ")
	return hand{tmp[:5]}, hand{tmp[5:]}
}

// reads given data file, returns all hands
// seperated into strings, which will hold
// hand of both players at a given time
func getAllHands(filepath string) []string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// determines winner between two hands,
// using clause 1 i.e. high card value
// we'll there might be a situation, when both
// hands having same high card i.e. `Q`,
// then we'll require to compare next high card value,
// if there's also a tie, we'll go for next high card value
// in this manner, we'll keep going on, until we obtain a solution
// with out a tie or we exhaust availble cards, whichever happens first
func determineWinnerwrtHighCard(h1 hand, h2 hand) int {
	hCard1, hCard2 := 0, 0
	to := 4
	for ; to >= 0 && hCard1 == hCard2; to-- {
		hCard1, hCard2 = h1.getHighCard(0, to), h2.getHighCard(0, to)
	}
	if to < 0 && hCard1 == hCard2 {
		return -1
	}
	if hCard1 > hCard2 {
		return 1
	}
	return 2
}

// given two hands, which are having a tie
// at a given `rank`, we'll try to resolute conflict;
// at the very end, we'll resort to finding high card value
// between those two hands
//
// as per given description in question, there's always
// a clear winner, so I'm keeping this in mind
// but in case of previous function, where I tried finding winner
// using high card value, I've kept one provision, which returns
// -1, to denote it was unable to find any clear winner
func tieBreaker(h1 hand, h2 hand, rank int) int {
	winner := 0
	switch rank {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 10:
		winner = determineWinnerwrtHighCard(h1, h2)
	case 2:
		if tmp1, tmp2 := h1.getOnePair(), h2.getOnePair(); tmp1 > tmp2 {
			winner = 1
		} else if tmp1 < tmp2 {
			winner = 2
		} else {
			winner = determineWinnerwrtHighCard(h1, h2)
		}
	case 4:
		if tmp1, tmp2 := h1.getThreeOfAKind(), h2.getThreeOfAKind(); tmp1 > tmp2 {
			winner = 1
		} else if tmp1 < tmp2 {
			winner = 2
		} else {
			winner = determineWinnerwrtHighCard(h1, h2)
		}
	case 7:
		three1, pair1 := h1.getFullHouse()
		three2, pair2 := h2.getFullHouse()
		if three1 > three2 {
			winner = 1
		} else if three1 < three2 {
			winner = 2
		} else {
			if pair1 > pair2 {
				winner = 1
			} else if pair1 < pair2 {
				winner = 2
			} else {
				winner = determineWinnerwrtHighCard(h1, h2)
			}
		}
	case 8:
		if tmp1, tmp2 := h1.getFourOfAKind(), h2.getFourOfAKind(); tmp1 > tmp2 {
			winner = 1
		} else if tmp1 < tmp2 {
			winner = 2
		} else {
			winner = determineWinnerwrtHighCard(h1, h2)
		}
	}
	return winner
}

// given a pair of hands ( of two competing players )
// first we'll sort them ascendically using card
// values ( numeric equivalent )
// then we'll try to compute rank for each of two hands,
// if they tie, we'll go for tie breaking, else we'll simply
// return winner ( as per who's having a larger rank )
func getWinner(h1 hand, h2 hand) int {
	h1.sort()
	h2.sort()
	if rank1, rank2 := h1.getRank(), h2.getRank(); rank1 > rank2 {
		return 1
	} else if rank1 == rank2 {
		return tieBreaker(h1, h2, rank1)
	}
	return 2
}

// PokerHands - Computes how many times player one wins
func PokerHands() int {
	count := 0
	hands := getAllHands("./p054_poker.txt")
	for _, v := range hands[:len(hands)-1] {
		if getWinner(getHandsOfTwoPlayers(v)) == 1 {
			count++
		}
	}
	return count
}
