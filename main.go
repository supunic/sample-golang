package main

import (
	"fmt"
)

func main()  {
	var guess int
	fmt.Print("your guess? ")
	fmt.Scanf("%v", &guess)
	fmt.Printf("Your guess is %v\n", guess)
}
