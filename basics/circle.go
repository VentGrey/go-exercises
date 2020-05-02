package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64;
	fmt.Println("Ingrese el radio del círculo a evaluar");
	fmt.Scan(&radius);

	fmt.Println("El área del círculo es: ", math.Pi * (radius * radius));
	fmt.Println("La circunferencia del círculo es: ", math.Pi * (2 * radius));
}
