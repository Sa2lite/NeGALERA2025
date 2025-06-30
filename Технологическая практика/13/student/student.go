package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Course    int
	AvgGrade  float64
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "отличник"
	case s.AvgGrade >= 3.5:
		return "хорошист"
	default:
		return "отчислен"
	}
}

func main() {
	student := Student{
		Name:      "Данила Трампов",
		BirthYear: 2000,
		Course:    3,
		AvgGrade:  4.8,
	}

	fmt.Printf(student.Name, student.GetAge(), student.GetStatus())
}
