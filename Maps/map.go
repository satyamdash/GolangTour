package main

import "fmt"

func printmap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {

	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	printmap(colors)
	delete(colors, "red")
	printmap(colors)

}
