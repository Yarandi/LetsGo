package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Yourr rate is: ", input)
	convertedInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if convertedInput > 5 || convertedInput < 1 {
		fmt.Println("you should rate between 1 to 5")
		return
	}else {
	    fmt.Println(convertedInput)
	}

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("your rate is: ", convertedInput + 1)
	}
	
}