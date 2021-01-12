package main

import (
	"fmt"
)

func main() {
	answer := 6
	var guess int

	fmt.Print("your guess? ")
	fmt.Scanf("%v", &guess)

	if answer == guess {
		fmt.Println("Bingo!")
	} else if answer > guess {
		fmt.Println("The answer is higher!")
	} else {
		fmt.Println("The answer is lower!")
	}
}
