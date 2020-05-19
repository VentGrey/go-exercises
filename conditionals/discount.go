package main

import "fmt"

func main() {
	var compra float64;

	fmt.Println("Ingrese el total de compra");
	fmt.Scan(&compra);

	if compra > 2500 {
		desc := compra * 0.08;
		fmt.Println("El total de compra es:", compra - desc);
	} else {
		fmt.Println("El total de compra es:", compra);
	}
}
