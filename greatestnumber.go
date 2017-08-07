package main

import "fmt"

func greatestnumber(a ...int) int {
	var greatest int
	for i := 0; i < len(a); i++ {
		if greatest < a[i] {
			greatest = a[i];
		}
	}
	return greatest
}

func main() {
	fmt.Println(greatestnumber(2, 3, 7, 111, 345))
}
