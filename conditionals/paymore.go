package main

import "fmt"

func main() {
	var sue float64;

	fmt.Println("Ingrese el sueldo del trabajor");
	fmt.Scan(&sue);

	if sue < 1000 {
		fmt.Println("El sueldo nuevo es:", sue * 1.15);
	} else {
		fmt.Println("El sueldo nuevo es:", sue * 1.12);
	}
}
