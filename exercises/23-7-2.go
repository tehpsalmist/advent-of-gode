package exercises

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"8thday.dev/aog/utils"
)

var joker = rune(74)

var JokerDeck = map[rune]int{
	65: 13,
	75: 12,
	81: 11,
	84: 9,
	57: 8,
	56: 7,
	55: 6,
	54: 5,
	53: 4,
	52: 3,
	51: 2,
	50: 1,
	74: 0,
}

type JokerHand struct {
	bid   int64
	cards string
	hand  []int
}

func NewJokerHand(h string) *JokerHand {
	hand := &JokerHand{}
	parts := strings.Fields(h)

	hand.bid, _ = strconv.ParseInt(parts[1], 0, 64)

	hand.cards = parts[0]

	handMap := map[rune]int{}
	jokers := 0
	for _, c := range hand.cards {
		if c == joker {
			jokers++
			continue
		}
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

	if jokers == 5 {
		hand.hand = append(hand.hand, 5)
	} else {
		hand.hand[0] += jokers
	}

	return hand
}

func Year23Day7Exercise2() {
	total := int64(0)

	input := utils.LineReaderString("./inputs/23-7.txt")

	allJokerHands := []JokerHand{}
	for _, h := range input {
		hand := NewJokerHand(h)
		allJokerHands = append(allJokerHands, *hand)
	}

	sort.Slice(allJokerHands, func(i, j int) bool {
		hand1 := allJokerHands[i]
		hand2 := allJokerHands[j]
		if hand1.hand[0] < hand2.hand[0] {
			return true
		}

		if hand1.hand[0] > hand2.hand[0] {
			return false
		}

		if len(hand1.hand) > 1 && hand1.hand[1] < hand2.hand[1] {
			return true
		}

		if len(hand1.hand) > 1 && hand1.hand[1] > hand2.hand[1] {
			return false
		}

		for i := 0; i < 5; i++ {
			if JokerDeck[rune(hand1.cards[i])] < JokerDeck[rune(hand2.cards[i])] {
				return true
			}

			if JokerDeck[rune(hand1.cards[i])] > JokerDeck[rune(hand2.cards[i])] {
				return false
			}
		}

		return false
	})

	for rank, hand := range allJokerHands {
		total += int64(rank+1) * hand.bid
	}

	fmt.Println(total)
}
