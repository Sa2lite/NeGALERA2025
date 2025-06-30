package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "Go")
	slice = append(slice, "Python")
	slice = append(slice, "JavaScript")

	for i, item := range slice {
		fmt.Printf("Индекс %d: %s\n", i, item)
	}

	slice = append(slice, "Java", "C++")
	fmt.Println(slice)
}
