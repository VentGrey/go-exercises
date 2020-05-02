package main

import "fmt"

func main() {
	var num1 float64;
	var num2 float64;

	fmt.Println("Ingrese el primer número")
	fmt.Scan(&num1)

	fmt.Println("Ingrese el segundo número")
	fmt.Scan(&num2)


	fmt.Println("La suma de los números es: ", num1 + num2)
	fmt.Println("La diferencia de los números es: ", num1 - num2)
	fmt.Println("El producto de los números es: ", num1 * num2)
	fmt.Println("La división de los números es: ", num1 / num2)
}
