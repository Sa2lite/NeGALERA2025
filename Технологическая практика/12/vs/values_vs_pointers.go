package main

import "fmt"

func incrementByValue(num int) {
	num += 10
	fmt.Println(num)
}

func incrementByPointer(num *int) {
	*num += 10
	fmt.Println(*num)
}

func main() {
	number := 5

	incrementByValue(number)
	fmt.Println(number)

	incrementByPointer(&number)
	fmt.Println(number)
}
