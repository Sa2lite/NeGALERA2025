package main

import "fmt"

func main() {
	var intVar int = 1
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 1 << 30
	var int64Var int64 = 1 << 62

	var uintVar uint = 42
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 1 << 31
	var uint64Var uint64 = 1 << 63

	var float32Var float32 = 3.14
	var float64Var float64 = 3.141592653589793

	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 3 + 4i

	var boolVar bool = true

	var stringVar string = "СИМОНОВ С. Е."

	fmt.Println("int:", intVar, "int8:", int8Var, "int16:", int16Var, "int32:", int32Var, "int64:", int64Var)
	fmt.Println("uint:", uintVar, "uint8:", uint8Var, "uint16:", uint16Var, "uint32:", uint32Var, "uint64:", uint64Var)
	fmt.Println("float32:", float32Var, "float64:", float64Var)
	fmt.Println("complex64:", complex64Var, "complex128:", complex128Var)
	fmt.Println("bool:", boolVar)
	fmt.Println("string:", stringVar)
}
