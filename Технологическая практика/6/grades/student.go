package main

import "fmt"

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Алексей", 5)
	addGrade(grades, "Мария", 4)
	addGrade(grades, "Иван", 3)

	fmt.Printf("Оценка Марии: %d\n", findGrade(grades, "Мария"))

	removeGrade(grades, "Иван")
	fmt.Println("После удаления", grades)
}

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func findGrade(grades map[string]int, name string) int {
	return grades[name]
}

func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
}
