package main

import "fmt"

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	languages := []string{"Go", "Python", "JavaScript", "Java", "C++"}
	fmt.Println("Исходный срез:", languages)

	languages = removeElement(languages, 2)
}
