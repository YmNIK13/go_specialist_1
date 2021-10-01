package main

import "fmt"

func main() {
	var a, b, c int64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("%d%d%d", c, b, a)
}
