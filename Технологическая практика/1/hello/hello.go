package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Егор"
	currentTime := time.Now().Format("02.01.2006")

	fmt.Println(name)
	fmt.Println(currentTime)
}
