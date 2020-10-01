package main

import (
	"fmt"
	"math"
)

func quad() {
	var (
		a float64
		b float64
		c float64

		item1 float64
		item2 float64
	)

	fmt.Print("What is a?: ")
	_, err_a := fmt.Scan(&a)
	fmt.Print("What is b?: ")
	_, err_b := fmt.Scan(&b)
	fmt.Print("What is c?: ")
	_, err_c := fmt.Scan(&c)

	item1 = ((-b) + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	item2 = ((-b) - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	_ = err_a
	_ = err_b
	_ = err_c

	fmt.Println("X = ", item1, " and X = ", item2)
}
