package main

import (
	"fmt"
)

func Sqrt(x float64) string {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	if x < 0 {
		return fmt.Sprintf("%fi", z)
	}
	return fmt.Sprintf("%f", z)
}

func main() {
	fmt.Println(Sqrt(2))
}
