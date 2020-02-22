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
// a mapping of card value to card suits
// if we've a multiple cards of same value which
// belong to different suits, then we'll keep them
// in a container, for that certain card value
//
// e.g. we've a hand of cards [5C 5D 9C AD AC]
// which can be represented as a map {3:[C D] 7:[C] 12:[D C]}
// where 3, 7, 12 i.e. keys are numeric representation of card values, see more about it below
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

func (h *hand) sort() {
	for i := 0; i < len(h.cards); i++ {
		for j := 0; j < len(h.cards)-i-1; j++ {
			if needsSwap(h.cards[j][:1], h.cards[j+1][:1]) {
				swap(&h.cards[j], &h.cards[j+1])
			}
		}
	}
}

func (h hand) getHighCard(from int, to int) int {
	max := 0
	for k := range h.getMap(from, to) {
		if k > max {
			max = k
		}
	}
	return max
}

func (h hand) getOnePair() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 2 && k > max {
			max = k
		}
	}
	return max
}

func (h hand) getTwoPairs() []int {
	pairs := []int{}
	for k, v := range h.getMap(0, 4) {
		if len(v) == 2 {
			pairs = append(pairs, k)
		}
	}
	if len(pairs) == 2 {
		return bubbleSort(pairs)
	}
	return nil
}

func (h hand) getThreeOfAKind() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 3 && k > max {
			max = k
		}
	}
	return max
}

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

func (h hand) getFullHouse() (int, int) {
	return h.getThreeOfAKind(), h.getOnePair()
}

func (h hand) getFourOfAKind() int {
	max := 0
	for k, v := range h.getMap(0, 4) {
		if len(v) == 4 && k > max {
			max = k
		}
	}
	return max
}

func (h hand) isStraightFlush() bool {
	return h.isStraight() && h.isFlush()
}

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

func cardValueToIndex(a string) int {
	cardVToIdxMap := map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
	return cardVToIdxMap[a]
}

func needsSwap(a string, b string) bool {
	if aIdx, bIdx := cardValueToIndex(a), cardValueToIndex(b); aIdx < bIdx {
		return false
	} else if aIdx == bIdx {
		return false
	}
	return true
}

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

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swapNum(&arr[j], &arr[j+1])
			}
		}
	}
	return arr
}

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

// PokerHands - ...
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
