package main

import "fmt"

func main() {
	// Массивы. Основы
	// 1. Определение массива
	// Создадим массив под хранение 5-ти целочисленных элементов
	var arr [5]int // при инициализации массива важно передать информацию - сколько элементов в нем будет
	fmt.Println("This is my array: ", arr)

	// 2. Определение элементов массива (после предварительной инициализации)
	// Необходимо обратиться к элементу массива через индекс, начинаются с 0
	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720
	fmt.Println("After element init: ", arr)

	// 3. Определение массива с указанием элементов массива
	// Если при инициализации количество элементов меньше номинальной длинны
	// то остальные элементы заполнятся нулями
	newArr := [7]int{10, 20, 30, 40, 50}
	fmt.Println("Short declaration and init: ", newArr)

	// 4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]: ", arrWithValues, " Length: ", len(arrWithValues))

	// 5. Массив - это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях - массив копируется (жестко на уровне компиляции)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First arr: ", first)
	fmt.Println("second arr: ", second)

	/*
		// 6. Массив и его размер - это две составляющие его типа (Размер и массив - часть типа)
		var aArr [5]int
		var bArr [6]int

		aArr[0] = 100
		// cannot use aArr (type [5]int) as type [6]int in assignment
		// bArr = aArr
		// */

	// 7. Итерирование по массиву
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d - %.1f\n", i, floatArr[i])
	}

	// 8. Итерирование по массиву через оператор range
	var sum float64
	for id, val := range floatArr {
		sum += val
		fmt.Printf("range - %d - %.1f\n", id, val)
	}

	fmt.Println("Sum total is: ", sum)

	// 9. Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}

	// 10. Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimensional array: ", words)

	// 11. Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}
}
