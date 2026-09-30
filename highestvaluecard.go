// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type resultVal struct {
	value int
	raw   string
}

var rankMap = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

var suitMap = map[string]int{
	"C": 1,
	"D": 2,
	"H": 3,
	"S": 4,
}

func parseCards(card string) resultVal {
	var rank string
	var suit string
	var RankValue int
	var SuitValue int
	if len(card) == 3 {
		rank = card[:2]
		suit = card[2:]
	} else {
		rank = card[0:1]
		suit = card[1:]
	}
	fmt.Println("Rank and suit", rank, suit)
	RankValue = rankMap[rank]
	SuitValue = suitMap[suit]
	fmt.Println("Rank and suit", RankValue, SuitValue)
	return resultVal{
		value: (RankValue * SuitValue),
		raw:   card,
	}

}

func highestCards(cards []string) string {
	highestValue := parseCards(cards[0])
	for i := 1; i < len(cards); i++ {
		curr := parseCards(cards[i])
		if highestValue.value < curr.value {
			highestValue = curr
		}
	}
	return highestValue.raw
}

func main() {
	inputcards := []string{"10S", "AD", "KC", "2S", "9C"}
	fmt.Println("Highest Value Card is", highestCards(inputcards))
	fmt.Println("Hello, 世界")
}
