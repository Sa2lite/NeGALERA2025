package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n",
		s.Name, s.Age, s.Course, s.AvgGrade)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func main() {
	student := Student{
		Name:     "Иван Петров",
		Age:      20,
		Course:   2,
		AvgGrade: 4.5,
	}

	student.PrintInfo()

	student.Promote()

	student.AvgGrade = 4.7
	fmt.Printf("Новый средний балл: %.2f\n", student.AvgGrade)
}
