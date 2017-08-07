package main

import "fmt"

func half(num int) (int, bool) {
	return num / 2, num%2 == 0
}
func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
