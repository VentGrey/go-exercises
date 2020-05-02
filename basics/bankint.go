package main

import "fmt"

func main() {
	var md float64;
	var tasa float64;

	fmt.Println("Ingrese la cantidad de dinero a guardar:");
	fmt.Scan(&md);

	fmt.Println("Ingrese la tasa de interés a cobrar:");
	fmt.Scan(&tasa);

	fmt.Println("Total:", md + (md * tasa));
}
