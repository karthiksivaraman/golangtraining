package main

import "fmt"

func main() {
	var small int32
	var large int32
	fmt.Print("Enter a small number:")
	fmt.Scan(&small)
	fmt.Print("Enter a large number:")
	fmt.Scan(&large)
	fmt.Printf("Remainder of Large divided by small is: %d", large%small)
}
