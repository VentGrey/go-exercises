package main

import (
	"fmt"
	"math"
)

func main() {
	var rad float64;

	fmt.Println("Ingrese el radio de su esfera");
	fmt.Scan(&rad);

	fmt.Println("Area:", 4 * math.Pi * rad);
	fmt.Println("Volumen: ", (4.0/3.0) * math.Pi * math.Pow(rad, 3));
}
