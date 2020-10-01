package main

import (
	"fmt"
)

func heatEqu() {
	var answer string
	var q float64
	var m float64
	var c float64
	var dt float64
	var fi float64
	var ini float64
	fmt.Print("What are we solving for? (q, m, c, t): ")
	_, err_ans := fmt.Scan(&answer)
	if answer == "q" {
		fmt.Print("What is m?: ")
		_, err_m := fmt.Scan(&m)
		fmt.Print("What is the final?: ")
		_, err_fi := fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		_, err_ini := fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		_, err_c := fmt.Scan(&c)
		q = m * c * (fi - ini)
		_ = err_ans
		_ = err_m
		_ = err_fi
		_ = err_ini
		_ = err_c
		fmt.Println("q = ", q)
	} else if answer == "c" {
		fmt.Print("What is q?: ")
		_, err_q := fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		_, err_fi := fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		_, err_ini := fmt.Scan(&ini)
		fmt.Print("What is m?: ")
		_, err_m := fmt.Scan(&m)
		c = q / ((fi - ini) * m)
		_ = err_ans
		_ = err_q
		_ = err_fi
		_ = err_ini
		_ = err_m
		fmt.Println("q = ", q)
	} else if answer == "m" {
		fmt.Print("What is q?: ")
		_, err_q := fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		_, err_fi := fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		_, err_ini := fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		_, err_c := fmt.Scan(&c)
		m = q / (c * (fi - ini))
		_ = err_ans
		_ = err_q
		_ = err_fi
		_ = err_ini
		_ = err_c
		fmt.Println("m = ", m)
	} else if answer == "t" {
		fmt.Print("What is m?: ")
		_, err_m := fmt.Scan(&m)
		fmt.Print("What is the final?: ")
		_, err_fi := fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		_, err_ini := fmt.Scan(&ini)
		fmt.Print("What is q?: ")
		_, err_q := fmt.Scan(&q)
		_ = err_ans
		_ = err_m
		_ = err_fi
		_ = err_ini
		_ = err_q
		dt = q / (c * m)
		fmt.Println("ΔT = ", dt)
	}

}
