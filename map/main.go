package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#fe45ab",
		"white": "#ffffff",
	}

	colors["yellow"] = "hello"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
