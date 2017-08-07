package main

import "fmt"

func main() {
	half := func(num int) (int, bool) {
		return num / 2, num%2 == 0;
	}
	fmt.Println(half(1))
	fmt.Println(half(2))
}
