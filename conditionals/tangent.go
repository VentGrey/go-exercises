package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64;

	fmt.Println("Ingrese el delta a calcular");
	fmt.Scan(&a);

	if ( math.Cos(a) != 0 ) {
		fmt.Println("La tangente de a es:", (math.Sin(a) / math.Cos(a)));
	} else if ( math.Sin(a) != 0 ) {
		fmt.Println("La cotangente de a es:", (math.Cos(a) / math.Sin(a)));
	} else {
		fmt.Println("No se puede calcular la tangente o cotangente");
	}
}
