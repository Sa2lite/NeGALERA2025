package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Scan(&num)

		if num >= 1 && num <= 10 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
