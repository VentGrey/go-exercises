package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64;

	fmt.Println("Ingrese el lado del hexaedro");
	fmt.Scan(&l);

	fmt.Println("AB:", math.Pow(l, 2));
	fmt.Println("AL:", 4 * math.Pow(l, 2));
	fmt.Println("AT:", 6 * math.Pow(l, 2));
	fmt.Println("Vol:", math.Pow(l, 3));
}
