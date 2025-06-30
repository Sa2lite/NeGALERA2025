package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E
	circleArea := pi * math.Pow(5, 2)
	expValue := math.Pow(e, 2)

	fmt.Println("Пи:", pi)
	fmt.Println("Экспонента:", e)
	fmt.Println("Площадь круга:", circleArea)
	fmt.Println("e^2 =", expValue)
}
