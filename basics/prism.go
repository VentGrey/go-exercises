package main

import "fmt"

func main() {
	var per float64;
	var apo float64;
	var alt float64;

	fmt.Println("Ingrese el perímetro del prisma");
	fmt.Scan(&per);

	fmt.Println("Ingrese el apotema del prisma");
	fmt.Scan(&apo);

	fmt.Println("Ingrese la altura del prisma");
	fmt.Scan(&alt);

	AB := (per * apo) / 2
	AL := per * alt

	fmt.Println("AB:", AB);
	fmt.Println("AL:", AL);
	fmt.Println("AT:", 2 * AB + AL);
	fmt.Println("Vol", AB * alt);
}
