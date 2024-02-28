package matriz

import "math/rand"

func GenerarMatrizCuadradaAleatoria(n int) [][]int {
	var matriz [][]int

	for i := 0; i < n; i++ {
		var fila []int
		for j := 0; j < n; j++ {
			numeroAleatorio := rand.Intn(100)
			fila = append(fila, numeroAleatorio)
		}
		matriz = append(matriz, fila)
	}

	return matriz
}
