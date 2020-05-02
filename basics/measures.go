package main

import "fmt"

func main() {
	var ECA float64;

	fmt.Println("Ingrese el número de acres que posee: ");
	fmt.Scan(&ECA);
	fmt.Println("Usted posee", ECA * 0.40 ,"hectáreas");
}
