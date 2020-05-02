package main

import "fmt"

func main() {
	var can float64;

	fmt.Println("Ingrese cuantos dólares posee: ");
	fmt.Scan(&can)

	fmt.Println("Usted posee", can * 18.20, "pesos Mexicanos.");
}
