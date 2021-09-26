package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type thaiBot struct {
}

type otherBot struct {
}

func main() {
	eb := englishBot{}
	tb := thaiBot{}
	ob := otherBot{}

	printGreeting(eb)
	printGreeting(tb)
	printGreeting(ob)
	printGreeting(nil)
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (thaiBot) getGreeting() string {
	return "Sawasdee"
}

func (otherBot) getGreeting() string {
	return "!#%$^@"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(tb thaiBot) {
// 	fmt.Println(tb.getGreeting())
// }
