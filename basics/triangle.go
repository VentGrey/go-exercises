package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64;
	var x2 float64;
	var x3 float64;
	var y1 float64;
	var y2 float64;
	var y3 float64;

	// Punto 1
	fmt.Println("Ingrese x1");
	fmt.Scan(&x1);
	fmt.Println("Ingrese y1");
	fmt.Scan(&y1);

	// Punto 2
	fmt.Println("Ingrese x2");
	fmt.Scan(&x2);
	fmt.Println("Ingrese y2");
	fmt.Scan(&y2);

	// Punto 3
	fmt.Println("Ingrese x3");
	fmt.Scan(&x3);
	fmt.Println("Ingrese y3");
	fmt.Scan(&y3);

	// Punto 1 - Punto 2
	D1 := math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2));
	// Punto 2 - Punto 3
	D2 := math.Sqrt(math.Pow(x2 - x3, 2) + math.Pow(y2 - y3));
	// Punto 3 - Punto 1
	D3 := math.Sqrt(math.Pow(x3 - x1, 2) + math.Pow(y3 - y1));

	Area := (1.0/2.0) * math.Abs(x1 * (y2-y3) + x2 * (y3-y1) + x3 (y1-y2));

	fmt.Println("Area:", Area);
	fmt.Println("Per:", D1 + D2 + D3);
}
