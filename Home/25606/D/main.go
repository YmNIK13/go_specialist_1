package main

import "fmt"

func main() {
	var name, lastName, age string
	fmt.Scan(&name)
	fmt.Scan(&lastName)
	fmt.Scan(&age)

	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %s . Студент BPS", name, lastName, age)
}
