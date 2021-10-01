package main

import "fmt"

func main() {
	var a, b int64

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print(a*a + 2*a*b + b*b)
}
