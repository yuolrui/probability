package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	ids  []int32
	rate float64
}

var cards []Card

func init() {
	card1 := Card{ //type1,rate 1%
		ids:  []int32{1, 2, 3, 8},
		rate: 0.01,
	}
	card2 := Card{ //type2,rate 5%
		ids:  []int32{201, 202, 203},
		rate: 0.05,
	}
	card3 := Card{
		ids:  []int32{301, 302, 303, 304, 306, 311, 312, 313, 333, 334},
		rate: 1 - card1.rate - card2.rate,
	}
	cards = []Card{card1, card2, card3}
}

func drawCard(times int) []int32 {
	result := make([]int32, times)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		rateSum := 0.0
		for _, card := range cards {
			rateSum += card.rate
			if rand.Float64() < rateSum {
				result[i] = card.ids[rand.Intn(len(card.ids))]
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(drawCard(100))
}
