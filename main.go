package main

import (
	"fmt"
)

func main()  {
	answer := 6
	var guess int

	fmt.Print("your guess? ")
	fmt.Scanf("%v", &guess)
	
	if answer == guess {
		fmt.Println("Bingo!")
	} else {
		fmt.Println("Boo...")
	}
}
