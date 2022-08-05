package main

import (
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"time"
)

func main() {
	shufflecards()
}

func poker(sCards [52]string) {
	err := exec.Command("clear").Run()
	if err != nil {
		return
	}
	var Deck []string
	var ChosenIndexes []int
	var isTrue = false
	var count int
	for {
		if count >= 3 {
			break
		}
		r := rand.Intn(len(sCards) - 1)
		if len(ChosenIndexes) > 0 {
			for i := len(ChosenIndexes) - 1; i > 0; i-- {
				if ChosenIndexes[i] == r {
					isTrue = true
					break
				}
			}
			if isTrue == true {
				continue
			}
		}
		ChosenIndexes = append(ChosenIndexes, r)
		Deck = append(Deck, sCards[r])
		count++
	}
	var CountOfPlayers int
	fmt.Println(Deck)
	for {
		fmt.Println("Please enter the number of players: ")
		_, err2 := fmt.Scanln(&CountOfPlayers)
		if err2 != nil {
			continue
		}
		if CountOfPlayers > 8 || CountOfPlayers < 3 {
			fmt.Println("Maximum number of players is 8 and the minimum is 3(including you)")
			CountOfPlayers = 0
		} else {
			break
		}
	}
	var countindex int
	for ; CountOfPlayers >= 0; CountOfPlayers-- {
		for i := len(ChosenIndexes) - 1; i > 0; i-- {
			l := ChosenIndexes[i]
			if sCards[countindex] == sCards[l] {
				continue
			}
		}
		countindex++
	}
}

func shufflecards() {
	var Cards = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var TypeOfCards = []string{"C", "S", "D", "H"}
	var ShuffledCards [52]string
	var count int
	var counttype int
	for i := (len(Cards) * len(TypeOfCards)) - 1; i >= 0; i-- {
		if math.Mod(float64(count), float64(len(Cards)-1)) == 0 && count != 1 && count != 0 {
			ShuffledCards[i] = Cards[count] + TypeOfCards[counttype]
			count = 0
			if counttype != len(TypeOfCards)-1 {
				counttype++
			} else {
				break
			}
			continue
		}
		ShuffledCards[i] = Cards[count] + TypeOfCards[counttype]
		count++

	}
	rand.Seed(time.Now().UnixNano())
	for i := len(ShuffledCards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ShuffledCards[i], ShuffledCards[j] = ShuffledCards[j], ShuffledCards[i]
	}
	poker(ShuffledCards)
}
