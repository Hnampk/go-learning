package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {

	for index, card := range d {

		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName) // byte slides or error

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cards := strings.Split(string(bs), ",") // slides of card

	return deck(cards)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	realRand := rand.New(source)

	for i := range d {
		randomPosition := realRand.Intn(len(d) - 1)

		d[i], d[randomPosition] = d[randomPosition], d[i]
	}
}
