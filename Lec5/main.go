package main

import (
	"fmt"
	"strings"
)

func main() {
	// Классический условный оператор
	// if condition {
	// 	// body
	// }

	// Условный оператор с блоком else
	// if condition {

	// } else {

	// }

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	// if condition {

	// } else if condition2 {

	// } else if condition2 {

	// } else {

	// }

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	//Инициализация в блоке условного оператора
	// Блок присваивания только :=
	// инициализированная переменная видна только в блоке условия
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
		fmt.Println(num)
	}

	// НЕ ИДУОМАТИЧНО
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}
	// Стоит избегать блока ELSE, связано с использованием go-рутин

	// Идиоматично
	if height := 100; height > 100 {
		fmt.Println("Height > 100")
		return
	}
	fmt.Println("Height <= 100")
}
