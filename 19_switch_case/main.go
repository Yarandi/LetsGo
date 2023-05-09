package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		//In Go, switch statements are case-oriented, and a case block will execute if the specified condition is matched.
		//The keyword fallthrough allows us to execute the next case block without checking its condition.
		//In other words, we can merge two case blocks by using the keyword fallthrough.
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again")
	default:
		fmt.Println("what was that!!")
	}

}
