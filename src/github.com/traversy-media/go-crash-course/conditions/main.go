package main

import "fmt"

func main() {
	x := 50
	y := 10

	// If else
	if x < y {
		fmt.Println("x is less than y")
	}else{
		fmt.Println("x is not less than y")
	}

	color := "green"
	// Switch
	switch color {
		case "red":
			fmt.Println("some text A")
		case "blue":
			fmt.Println("some text B")
		default:
			fmt.Println("default text")
	}
}