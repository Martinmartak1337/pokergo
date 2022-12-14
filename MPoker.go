package main

import (
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello this is the Mpoker binary/application!\n")
	time.Sleep(time.Millisecond * 500)
	ContinueTheGame(HandSAndFlop(ShuffleCards()))
}

func ContinueTheGame(a int, b int, c int, d int, e [46]string, f [3]string, g [2]string) {
	var smallblindbet int
	var bigblindbet int
	switch d {
	case b:
		for {
			fmt.Println("Enter a 1USD bet!\n")
			_, bet1 := fmt.Scanln(&smallblindbet)
			if bet1 != nil {
				continue
			}
			if int(smallblindbet) != 1 {
				continue
			} else {
				break
			}
		}
	case c:
		for {
			fmt.Println("Enter a 2USD bet!\n")
			_, bet1 := fmt.Scanln(&bigblindbet)
			if bet1 != nil {
				continue
			}
			if int(bigblindbet) != 2 {
				continue
			} else {
				break
			}
		}
	case a:
		break
	default:
		fmt.Printf("Your position is \n", d)
		break
	}

	fmt.Println("\nThis is the flop: \n", f)
}

func HandSAndFlop(p *[52]string) (int, int, int, int, [46]string, [3]string, [2]string) {
	err := exec.Command("clear").Run()
	if err != nil {
		goto loc
	}
loc:
	var indexofcard int
	var CountOfPlayers int
	for {
		fmt.Println("Please enter the number of players: ")
		_, err2 := fmt.Scanln(&CountOfPlayers)
		if err2 != nil {
			continue
		}
		if CountOfPlayers > 8 || CountOfPlayers < 3 {
			fmt.Println("Maximum number of players is 8 and the minimum is 3 (including you).")
			CountOfPlayers = 0
		} else {
			break
		}
	}
	memCount := CountOfPlayers
	rand.Seed(time.Now().UnixNano())
	var DealersPosition = rand.Intn(CountOfPlayers)
	var flop [3]string
	for i := 0; i <= 2; i++ {
		flop[i] = p[indexofcard]
		indexofcard++
	}
	var InvCards [46]string
	InvCards[0] = ""
	var count = 1
	var myhand [2]string
	CountOfPlayers *= 2
	for ; CountOfPlayers > 0; CountOfPlayers-- {
		InvCards[count] = p[indexofcard]
		count++
		indexofcard++
		continue
	}
	for j := 1; j >= 0; j-- {
		myhand[j] = p[indexofcard]
		indexofcard++
	}
	var SmallBlind int
	var BigBlind int
	if DealersPosition == 0 {
		DealersPosition = 1
		SmallBlind = memCount - 1
		BigBlind = memCount - 2
	} else {
		SmallBlind = memCount - 1
		BigBlind = memCount - 2
	}
	rand.Seed(time.Now().UnixNano())
	var PlayersPosition = rand.Intn(memCount)
	fmt.Printf("This is your hand:\n %s ", myhand)
	fmt.Printf("\n")
	return DealersPosition, SmallBlind, BigBlind, PlayersPosition, InvCards, flop, myhand
}
func ShuffleCards() *[52]string {
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
	return &ShuffledCards
}
