package main

import "fmt"

func main() {
	var mon float64;

	fmt.Println("Ingrese el monto total del vehículo");
	fmt.Scan(&mon);

	fmt.Println("El enganche del vehículo es:", mon * 0.35);
	fmt.Println("Las mensualidades del vehículo se pagarán en montos de:",
		(mon - (mon * 0.35)) / 18);
}
