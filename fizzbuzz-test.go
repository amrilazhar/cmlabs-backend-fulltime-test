package main

import (
	"fmt"
)

func main() {
	count := 100
	for i := 1; i <= count; i++ {
		tempString := ""
		if i%3 == 0 {
			tempString += "Fizz"
		}
		if i%5 == 0 {
			tempString += "Buzz"
		}

		if tempString == "" {
			fmt.Println(i)
		} else {
			fmt.Println(tempString)
		}
	}

}
