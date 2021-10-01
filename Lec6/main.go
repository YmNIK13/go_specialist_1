package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		for init; condition; post {
			init - блок инициализации
			condition - условие завершения
			post - изменение переменной цикла
		}
	*/

	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}

	// Важный момент: В качестве init может быть использовано ТОЛЬКО :=

	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Break value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Continue value: %d\n", i)
	}
	fmt.Println("After for loop with Continue")

	//вложенный цикл
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идее треугольник")

outer:
	// Иногда выйти с двух циклов, можно с помошью Lable
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)

			if i == j {
				break outer
			}
		}
		fmt.Println()
	}

	fmt.Println("Переход после break outer")

	// модификации цикла for
	// 1. Классический цикл while do
	var loopVar int = 0
	/*
		while (loopVar < 10){
			...
			loopVar++
		}
	*/
	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}

	// 2. классический бесконечный цикл
	var password string
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)

		if !strings.Contains(password, "1234") {
			fmt.Println("Password accepted")
			break
		}

		fmt.Sprintln("Weak password. Try again")
	}

	// 3. Цикл с множественными перменными цикла
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}
}
