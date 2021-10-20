package main

import (
	"fmt"
)

func main() {
	var (
		drink   string
		meal    string
		subMeal string
		time    string
	)

	fmt.Scan(&drink, &meal, &subMeal, &time)
	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)
}
