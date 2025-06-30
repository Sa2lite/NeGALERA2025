package main

import "fmt"

func main() {
	for num := 2; num <= 100; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			for num%i == 0 {
				isPrime = false
				break
			}
		}

		for isPrime {
			fmt.Print(num, " ")
			break
		}
	}
	fmt.Println()
}
