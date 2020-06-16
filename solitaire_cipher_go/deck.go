package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cardInfo struct {
	name  string
	value int
}

type deck []cardInfo

func newDeck() deck {
	cards := deck{}

	suits := []string{"Clubs", "Diamonds", "Spades", "Hearts"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "Jack", "Queen", "King"}

	var i int = 1
	for _, suit := range suits {
		for _, value := range values {
			info := value + " of " + suit
			cards = append(cards, cardInfo{info, i})
			i++
		}
	}
	cards = append(cards, cardInfo{"Joker A", 53}, cardInfo{"Joker B", 53})

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func (d deck) saveToFile(filename string) error {
// 	//WriteFile returns an error if one occurs
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }

// func deckFromFile(filename string) deck {
// 	bytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		fmt.Println("Quitting application...")
// 		os.Exit(1)
// 	}

// 	s := strings.Split(string(bytes), ",")
// 	return deck(s)
// }

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		//swap the current index and a random position
		randI := rand.Intn(len(d) - 1)
		d[i], d[randI] = d[randI], d[i]
	}
}

func (d deck) generateKeystream() {
	//The deck in it's current order is considered the "key" here

}

func (d deck) findValue(value string) int {
	for i, v := range d {
		if v.name == value {
			return i
		}
	}
	return -1
}

func (d deck) moveDown(index int, positions int) {
	//add check for positions being less than 1 or greater than 53
	currentIndex := index
	swapIndex := index + 1

	for i := positions; i > 0; i-- {
		if swapIndex == len(d) {
			temp := d[currentIndex]
			for i := len(d) - 1; i > 0; i-- {
				d[i] = d[i-1]
			}
			d[0] = temp
		} else {
			d[currentIndex], d[swapIndex] = d[swapIndex], d[currentIndex]
		}
		currentIndex++
		if currentIndex == len(d) {
			currentIndex = 0
		}
		swapIndex = currentIndex + 1
	}
}
func (d deck) tripleCut() deck {
	jokerA := d.findValue("Joker A")
	jokerB := d.findValue("Joker B")

	var cut1 deck
	var cut2 deck
	var cut3 deck
	if jokerA < jokerB {
		cut1 = d[:jokerA]
		cut2 = d[jokerA : jokerB+1]
		cut3 = d[jokerB+1:]
	} else {
		cut1 = d[:jokerB]
		cut2 = d[jokerB : jokerA+1]
		cut3 = d[jokerA+1:]
	}
	var c deck
	c = append(cut3, cut2...)
	c = append(c, cut1...)
	return c
}

func (d deck) countCut() deck {
	bottomCard := d[len(d)-1]
	bottomCut := d[:bottomCard.value]
	midCut := d[bottomCard.value : len(d)-1]

	var c deck
	c = append(bottomCut, midCut...)
	c = append(c, bottomCard)

	return c
}

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}