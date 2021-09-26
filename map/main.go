package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#FF5733",
		"blue":   "#3380FF",
		"green":  "#33FF52",
		"yellow": "#FFEC33",
		"pink":   "#FF33E3",
		"orange": "#FF7A33",
	}

	printMap(colors)
	// for key, value := range colors {
	// 	fmt.Printf("%v: %v \n", key, value)
	// }

	car := make(map[int]int, 5)
	car[100] = 1000000
	// result, ok := car["toyota"]

}

func printMap(mapData map[string]string) {
	for key, value := range mapData {
		fmt.Println("Hex code for", key, "is", value)
	}
}
