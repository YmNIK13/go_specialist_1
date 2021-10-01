package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Boolean
	var firstBool bool
	fmt.Println(firstBool)

	// Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR: ", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	// Numerics. Integers
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)

	// Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	// Узнаем сколько байт занимает
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	// Эксперемент. При использовании коротноко объявления - тип дял целого числа платформозависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	// Эксперемент 2.
	var first32 int32 = 12
	var second64 int64 = 13

	// Используйте явное приведение типов
	fmt.Println(int64(first32) + (second64))

	// Эксперемент 3. Если проводятся арифметические операции над int и intX
	// надо приводить к одному типу
	var third64 int64 = 16123414
	var fourthInt int = 1231244
	fmt.Println(third64 + int64(fourthInt))
	// оперции с целыми числами + - * / % - остаток от деления
	// аналогичным образом устроены  uint8, uint16, uint32, uint64, uint

	// Numerics. Float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)

	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

	// Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// String
	name := "Федя"
	lastName := "Pupkin"
	concat := name + " " + lastName
	fmt.Println("Full name: ", concat)

	// Функция len возвращает количество элементов в наборе, тоесть байт со строкой
	fmt.Println("Length os string: ", name, len(name))

	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))
	// Rune - руна это один utf-ный символ.

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "CDE"

	fmt.Println(strings.Contains(totalString, subString))

	// rune -> alias int32, по сути код символа в UTF8
	var sampleRune rune
	var anotherRune rune = 'Q'
	var third rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", third)

	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "a"))

	// byte -> alias uint8
}
