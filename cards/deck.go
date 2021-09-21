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

// type card string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// return strings.Join((d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resultFromReadFile := string(bs)
	var arr []string = strings.Split(resultFromReadFile, ",")
	return deck(arr)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Intn(len(d) - 1))

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10-1) + 1)
	// for i := range d {
	// 	newPosition := rand.Intn(len(d) - 1)
	// 	d[i], d[newPosition] = d[newPosition], d[i]
	// }

	for i := range d {
		newPosition := r.Intn((len(d) - 1))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
