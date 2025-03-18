package main

import "fmt"

func conversor(k float64) float64 {
	return k - 273
}
func main() {
	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scan(&kelvin)
	celsius := conversor(kelvin)
	fmt.Printf("A temperatura em Celsius é: %.2f", celsius)
	fmt.Println("execução finalizada")
}
