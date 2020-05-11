package main

import (
	"fmt"
	"math"
)

func main() {
	var rad float64;
	var alt float64;
	var gene float64;

	fmt.Println("Ingrese el radio del cono");
	fmt.Scan(&rad);

	fmt.Println("Ingrese la altura del cono");
	fmt.Scan(&alt);

	fmt.Println("Ingrese la generatriz del cono");
	fmt.Scan(&gene);

	AB := math.Pi * (rad * rad);
	AL := math.Pi * rad * gene;

	fmt.Println("AB:", AB);
	fmt.Println("AL:", AL);
	fmt.Println("AT:", AB + AL);
	fmt.Println("Vol:", (1.0/3.0) * AB * alt);
}
