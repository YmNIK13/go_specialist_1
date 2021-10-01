package main

import "fmt"

func main() {
	var a, b int64

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("Периметр: %d\n", (a+b)*2)
	fmt.Printf("Площадь: %d", (a * b))
}
