package main

import "fmt"

func main() {
	var word1, word2, word3, word4, author string

	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Scan(&word3)
	fmt.Scan(&word4)
	fmt.Scan(&author)

	fmt.Printf("%s - %s\n", word4, author)
	fmt.Printf("%s - %s\n", word3, author)
	fmt.Printf("%s - %s\n", word2, author)
	fmt.Printf("%s - %s", word1, author)
}
