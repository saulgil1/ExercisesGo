package sumatorias

func SumaTotal(lista []int) int {
	// Se inicializa la variable
	sumaTotal := 0

	// Se realiza el ciclo for para recorrer la lista y dar el valor de la suma
	for i := 0; i < len(lista); i++ {
		sumaTotal += lista[i]
	}
	return sumaTotal
}
