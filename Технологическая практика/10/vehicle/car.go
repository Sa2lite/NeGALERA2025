package main

import "fmt"

type Engine struct {
	Power    int
	FuelType string
}

type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

func main() {
	car := Car{
		Brand: "БМД",
		Model: "4",
		Year:  2023,
		Engine: Engine{
			Power:    777,
			FuelType: "Дизель",
		},
	}

	fmt.Printf("%s %s %d года\n", car.Brand, car.Model, car.Year)
	fmt.Printf("Двигатель: %d л.с., тип топлива: %s\n", car.Engine.Power, car.Engine.FuelType)

}
