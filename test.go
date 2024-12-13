package main

import "fmt"

func main() {
	var x, y int

	fmt.Printf("Entre 2 nombres> ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Printf("Les nombres entrés sont %d et %d\n", x, y)
}
