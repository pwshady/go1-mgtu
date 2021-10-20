package main

import (
	"fmt"
)

func main() {
	var (
		first_name string
		last_name  string
		age        int
	)
	fmt.Scan(&first_name, &last_name, &age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", first_name, last_name, age)
}
