package main

import "fmt"

func main() {
	var price int
	fmt.Scan(&price)

	if price == 100 {
		fmt.Println("if 100")
	} else if price == 110 {
		fmt.Println("if 110")
	} else if price == 120 {
		fmt.Println("if 120")
	}

	// в switch-case запрещени дублирующие условия
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")

	default:
		// отрабатывает только в том случае если не один
		// из выше перечисленных кейсов не отработал
		fmt.Println("Default case")

	}

	// Case с множеством вариантов
	var ageGroup string = "b" // возрастные группы: "a", "b", "c", "d", "e"

	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("Age group 10-40")
	case "d", "e":
		fmt.Println("Age group 50-70")

	default:
		fmt.Println("You are too yong/old")

	}

	// Case с выражением
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too young")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 20 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you")
	}

	// Case с проваливаниями. Проваливание выполняю даже ложные кейсы
	var number int
	fmt.Scan(&number)

outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d is less then 200\n", number)
	case number > 1000:
		fmt.Printf("%d GREATER then 1000\n", number)
		if number%3 == 0 {
			break outer
		}
		fallthrough

	default:
		fmt.Printf("%d default\n", number)

	}

	// Гадость с терминацией for из switch
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}

	fmt.Println("END")
}
