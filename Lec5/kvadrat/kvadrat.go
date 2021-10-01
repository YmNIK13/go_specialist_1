package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите аргументы квадратного уравнения:")
	fmt.Println("a*x^2 + b*x + c = 0")
	var a, b, c float64

	fmt.Print("Введите A: ")
	fmt.Scan(&a)

	fmt.Print("Введите B: ")
	fmt.Scan(&b)

	fmt.Print("Введите C: ")
	fmt.Scan(&c)

	D := b*b - 4*a*c

	fmt.Printf("\nВычисляем дискриминант: D = %.1f*%.1f-4*%.1f*%.1f = %.1f\n", b, b, a, c, D)

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("X1 = %.3f  X2 = %.3f\n", x1, x2)
	} else if D == 0 {
		x12 := -b / (2 * a)
		fmt.Printf("X1 = X2 = %.3f\n", x12)
	} else {
		fmt.Println("Нет решений для действительных чисел")
	}

}
