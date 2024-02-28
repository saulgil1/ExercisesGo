package sumatorias

func SumaFilas(matriz [][]int) []int {
	// Se inicializa una lista del tamaño de la matriz
	sumaFilas := make([]int, len(matriz))

	// Se realiza el ciclo for para recorrer la matriz y sumarla de forma horizontal
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			sumaFilas[i] += matriz[i][j]
		}
	}
	return sumaFilas
}
