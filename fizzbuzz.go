package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		printNum := true
		if i%3 == 0 {
			fmt.Print("Fizz")
			printNum = false
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			printNum = false
		}
		if printNum {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
