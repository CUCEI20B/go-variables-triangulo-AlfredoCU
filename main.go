package main

// Importando nuestras bibliotecas.
import "fmt"

// Creamos una funcion principal.
func main() {
	// Declaramos variables de tipo float64.
	var (
		base   uint64
		height uint64
	)
	fmt.Scan(&base) // Tomamos de entrada un valor numérico.
	fmt.Scan(&height) // Tomamos de entrada un valor numérico.

	// Formula para sacar el área de un triángulo.
	res := (base * height) / 2
	fmt.Println(res)
}