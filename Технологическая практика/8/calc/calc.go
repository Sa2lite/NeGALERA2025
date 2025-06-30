package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Scan(&a)

	fmt.Print("Введите знак: ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	default:
		fmt.Println("FF")
	}
}
