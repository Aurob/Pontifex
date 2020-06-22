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

func (d deck) deal(handSize int) (deck, deck) {
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

func (d *deck) generateKeystream(c int) string {
	var keystream string
	for i := 0; i < c; i++ {
		jokerAPos := (*d).findValue("Joker A")
		(*d).moveUp(jokerAPos, 1)
		jokerBPos := (*d).findValue("Joker B")
		(*d).moveUp(jokerBPos, 2)

		(*d).tripleCut()
		(*d).countCut()

		outputCard := (*d).getOutputCard()
		if outputCard.value == 53 {
			var ks string
			ks = (*d).generateKeystream(1)
			keystream += ks
		} else {
			keystream += string(alpha[outputCard.value%26])
		}
	}
	return keystream
}

func (d deck) findValue(value string) int {
	for i, v := range d {
		if v.name == value {
			return i
		}
	}
	return -1
}

func (d deck) moveUp(index int, positions int) {
	//add check for positions being less than 1 or greater than 53
	currentIndex := index
	swapIndex := index + 1

	for i := positions; i > 0; i-- {
		if currentIndex == len(d)-1 {
			temp := d[currentIndex]
			for i := len(d) - 1; i > 1; i-- {
				d[i] = d[i-1]
			}
			d[1] = temp
			currentIndex = 0
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
func (d *deck) tripleCut() {
	jokerA := d.findValue("Joker A")
	jokerB := d.findValue("Joker B")

	var temp deck
	var firstJoker, secondJoker int
	if jokerA < jokerB {
		firstJoker = jokerA
		secondJoker = jokerB
	} else {
		firstJoker = jokerB
		secondJoker = jokerA
	}
	temp = append(temp, (*d)[secondJoker+1:]...)
	temp = append(temp, (*d)[firstJoker:secondJoker+1]...)
	temp = append(temp, (*d)[:firstJoker]...)
	(*d) = temp

}

func (d *deck) countCut() {
	bottomCard := (*d)[len((*d))-1]
	if bottomCard.value != 53 {
		bottomCut := (*d)[:bottomCard.value]
		midCut := (*d)[bottomCard.value : len((*d))-1]

		var temp deck
		temp = append(midCut, bottomCut...)
		temp = append(temp, bottomCard)

		(*d) = temp
	}
}

func (d *deck) getOutputCard() cardInfo {
	var outputCard cardInfo
	topCard := (*d)[0]
	if topCard.value == 53 {
		outputCard = (*d)[len((*d))-1]
	} else {
		outputCard = (*d)[topCard.value]
	}

	return outputCard
}

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
