package main

import (
	"fmt"
	"math"
)

func main() {
	// простейший вывод. Это вывод аргумента + '\n'
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")

	// функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")

	// форматированный вывод Printf - стандартный вывод os.Stdout
	// с флагами форматирования
	fmt.Printf("\n\nHello, my name is %s\nMy age is %d\n", "Bob", 42)

	//////////////////////////
	//////////////////////////
	// декларирование пепеременных
	var age int
	fmt.Println("My age is:", age)

	age = 32
	fmt.Println("Age after assignment:", age)

	// Декларирование и инициализация пользовательских переменных
	var height int = 183
	fmt.Println("My height is:", height)

	// в чем полустроговть типизации
	var uid = 123

	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	// Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar float32 = 20, 39.5
	fmt.Printf("FirstVar: %f, SecondVar: %f\n", firstVar, secondVar)

	fmt.Printf("FirstVar: %T, SecondVar: %T\n", firstVar, secondVar)

	var (
		personName string = "Bob"
		personAge         = 42
		personUid  int
	)

	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUid)

	// Немного странного
	var firstVar2, secondVar2 = "Vasja", 39.5
	fmt.Printf("firstVar2: %s, firstVar2: %f\n", firstVar2, secondVar2)
	fmt.Printf("firstVar2: %T, firstVar2: %T\n", firstVar2, secondVar2)

	// Немного хорошего.
	// Повторное декларирование переменной приводит к ошибке
	// var firstVar2 = 200

	// Короткое объявление
	count := 10
	fmt.Println("count: ", count)
	count++
	fmt.Println("count: ", count)

	// Множественно присваивание через :=
	aArg, bArg := 10, "Petja"
	fmt.Println(aArg, bArg)

	aArg, bArg = 30, "Kolja"
	fmt.Println(aArg, bArg)

	// Исключение из присваивание через :=

	aArg, bArg, cArg := 300, "Masha", 500
	fmt.Println(aArg, bArg, cArg)

	// Пример
	width, length := 20.7, 30.2
	fmt.Printf("Min dimensional of rectagle is %.3f \n", math.Min(width, length))
}
