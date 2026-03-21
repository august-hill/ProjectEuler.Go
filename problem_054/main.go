// Problem 54: Poker Hands
// How many hands does Player 1 win?
// Answer: 376

package main

import (
	_ "embed"
	"sort"
	"strings"

	"github.com/august-hill/ProjectEuler.Go/bench"
)

//go:embed poker.txt
var pokerData string

// HandRank represents poker hand rankings (higher = better)
type HandRank int

const (
	HighCard HandRank = iota + 1
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// Card represents a playing card
type Card struct {
	Value int // 2-14 (14 = Ace)
	Suit  byte
}

// Hand represents a poker hand with its classification
type Hand struct {
	Cards      []Card
	Rank       HandRank
	RankValues []int // Values for tiebreaking, in priority order
}

func parseCard(s string) Card {
	valueMap := map[byte]int{
		'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}
	return Card{
		Value: valueMap[s[0]],
		Suit:  s[1],
	}
}

func parseHand(cards []string) Hand {
	hand := Hand{Cards: make([]Card, 5)}
	for i, c := range cards {
		hand.Cards[i] = parseCard(c)
	}
	hand.detectRank()
	return hand
}

func (h *Hand) detectRank() {
	// Sort cards by value descending for easier processing
	sort.Slice(h.Cards, func(i, j int) bool {
		return h.Cards[i].Value > h.Cards[j].Value
	})

	// Count value frequencies
	freq := make(map[int]int)
	for _, c := range h.Cards {
		freq[c.Value]++
	}

	// Check for flush
	isFlush := true
	firstSuit := h.Cards[0].Suit
	for _, c := range h.Cards[1:] {
		if c.Suit != firstSuit {
			isFlush = false
			break
		}
	}

	// Check for straight (cards are sorted descending)
	isStraight := true
	for i := 0; i < 4; i++ {
		if h.Cards[i].Value-h.Cards[i+1].Value != 1 {
			isStraight = false
			break
		}
	}

	// Categorize by frequency
	var fours, trips, pairs, singles []int
	for val, count := range freq {
		switch count {
		case 4:
			fours = append(fours, val)
		case 3:
			trips = append(trips, val)
		case 2:
			pairs = append(pairs, val)
		case 1:
			singles = append(singles, val)
		}
	}

	// Sort each group descending
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))
	sort.Sort(sort.Reverse(sort.IntSlice(singles)))

	// Determine hand rank and tiebreaker values
	if isStraight && isFlush {
		if h.Cards[0].Value == 14 { // Ace high
			h.Rank = RoyalFlush
			h.RankValues = []int{14}
		} else {
			h.Rank = StraightFlush
			h.RankValues = []int{h.Cards[0].Value} // High card of straight
		}
	} else if len(fours) == 1 {
		h.Rank = FourOfAKind
		h.RankValues = append(fours, singles...)
	} else if len(trips) == 1 && len(pairs) == 1 {
		h.Rank = FullHouse
		h.RankValues = append(trips, pairs...)
	} else if isFlush {
		h.Rank = Flush
		h.RankValues = []int{h.Cards[0].Value, h.Cards[1].Value, h.Cards[2].Value, h.Cards[3].Value, h.Cards[4].Value}
	} else if isStraight {
		h.Rank = Straight
		h.RankValues = []int{h.Cards[0].Value}
	} else if len(trips) == 1 {
		h.Rank = ThreeOfAKind
		h.RankValues = append(trips, singles...)
	} else if len(pairs) == 2 {
		h.Rank = TwoPairs
		h.RankValues = append(pairs, singles...)
	} else if len(pairs) == 1 {
		h.Rank = OnePair
		h.RankValues = append(pairs, singles...)
	} else {
		h.Rank = HighCard
		h.RankValues = singles
	}
}

// Compare returns true if h beats other
func (h *Hand) Compare(other *Hand) bool {
	if h.Rank != other.Rank {
		return h.Rank > other.Rank
	}
	// Same rank - compare tiebreaker values lexicographically
	for i := 0; i < len(h.RankValues) && i < len(other.RankValues); i++ {
		if h.RankValues[i] != other.RankValues[i] {
			return h.RankValues[i] > other.RankValues[i]
		}
	}
	return false // Tie
}

func solve() int64 {
	lines := strings.Split(strings.TrimSpace(pokerData), "\n")
	player1Wins := 0

	for _, line := range lines {
		cards := strings.Fields(line)
		hand1 := parseHand(cards[0:5])
		hand2 := parseHand(cards[5:10])

		if hand1.Compare(&hand2) {
			player1Wins++
		}
	}

	return int64(player1Wins)
}

func main() { bench.Run(54, solve) }
