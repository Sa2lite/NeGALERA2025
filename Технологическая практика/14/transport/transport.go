package main

import "fmt"

type Transport interface {
	Move() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Move() string {
	return fmt.Sprintf("%s едет", c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("%s остановился", c.Model)
}

type Bus struct {
	RouteNumber int
}

func (b Bus) Move() string {
	return fmt.Sprintf("Автобус №%d едет", b.RouteNumber)
}

func (b Bus) Stop() string {
	return fmt.Sprintf("Автобус №%d остановился на остановке", b.RouteNumber)
}

type Tram struct {
	LineNumber int
}

func (t Tram) Move() string {
	return fmt.Sprintf("Трамвай №%d движется", t.LineNumber)
}

func (t Tram) Stop() string {
	return fmt.Sprintf("Трамвай №%d остановился на остановке", t.LineNumber)
}

func testTransport(t Transport) {
	fmt.Println(t.Move())
	fmt.Println(t.Stop())
}

func main() {
	car := Car{Model: "LADA 2107"}
	bus := Bus{RouteNumber: 42}
	tram := Tram{LineNumber: 5}

	testTransport(car)
	fmt.Println("---")
	testTransport(bus)
	fmt.Println("---")
	testTransport(tram)
}
