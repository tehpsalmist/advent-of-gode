package exercises

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"8thday.dev/aog/utils"
)

var Deck = map[rune]int{
	65: 13,
	75: 12,
	81: 11,
	74: 10,
	84: 9,
	57: 8,
	56: 7,
	55: 6,
	54: 5,
	53: 4,
	52: 3,
	51: 2,
	50: 1,
}

type Hand struct {
	bid   int64
	cards string
	hand  []int
}

func NewHand(h string) *Hand {
	hand := &Hand{}
	parts := strings.Fields(h)

	hand.bid, _ = strconv.ParseInt(parts[1], 0, 64)

	hand.cards = parts[0]

	handMap := map[rune]int{}
	for _, c := range hand.cards {
		if _, ok := handMap[c]; ok {
			handMap[c]++
		} else {
			handMap[c] = 1
		}
	}

	hand.hand = []int{}
	for _, l := range handMap {
		hand.hand = append(hand.hand, l)
	}

	sort.Slice(hand.hand, func(i, j int) bool { return hand.hand[i] > hand.hand[j] })

	return hand
}

func Year23Day7Exercise1() {
	total := int64(0)

	input := utils.LineReaderString("./inputs/23-7.txt")

	allHands := []Hand{}
	for _, h := range input {
		hand := NewHand(h)
		allHands = append(allHands, *hand)
	}

	sort.Slice(allHands, func(i, j int) bool {
		hand1 := allHands[i]
		hand2 := allHands[j]
		if hand1.hand[0] < hand2.hand[0] {
			return true
		}

		if hand1.hand[0] > hand2.hand[0] {
			return false
		}

		if hand1.hand[1] < hand2.hand[1] {
			return true
		}

		if hand1.hand[1] > hand2.hand[1] {
			return false
		}

		for i := 0; i < 5; i++ {
			if Deck[rune(hand1.cards[i])] < Deck[rune(hand2.cards[i])] {
				return true
			}

			if Deck[rune(hand1.cards[i])] > Deck[rune(hand2.cards[i])] {
				return false
			}
		}

		return false
	})

	for rank, hand := range allHands {
		if 790 < rank && rank < 805 {
			fmt.Println(hand)
		}
		total += int64(rank+1) * hand.bid
	}

	fmt.Println(total)
}
