package sumatorias

func SumaColumnas(matriz [][]int) []int {
	// Se inicializa una lista del tamaño de la matriz
	sumasColumnas := make([]int, len(matriz))

	// Se realiza el ciclo for para recorrer la matriz y sumarla de forma vertical
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			sumasColumnas[j] += matriz[i][j]
		}
	}
	return sumasColumnas
}
