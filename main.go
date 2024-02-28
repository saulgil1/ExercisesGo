package main

import (
	"fmt"

	"github.com/saulgil1/ExercisesGo/matriz"
	"github.com/saulgil1/ExercisesGo/sumatorias"
)

func main() {
	// Se genera la matriz aleatoria, con n filas y columnas
	matriz := matriz.GenerarMatrizCuadradaAleatoria(3)

	// Se calcula la suma
	sumRows := sumatorias.SumaFilas(matriz)
	sumColumns := sumatorias.SumaColumnas(matriz)
	totalSumRows := sumatorias.SumaTotal(sumRows)
	totalSumColumns := sumatorias.SumaTotal(sumColumns)

	// Se muestra la matriz por pantalla
	fmt.Println(matriz)

	// Se muestra la suma de las filas y vertical
	fmt.Println("La Suma de las filas es:", sumRows)
	fmt.Println("La Suma de las columnas es:", sumColumns)

	// Se muesta la suma Total de las filas y columnas
	fmt.Println("La sumatoria de las sumas de las filas es:", totalSumRows)
	fmt.Println("La sumatoria de las sumas de las columnas es:", totalSumColumns)

	// Se muestra por pantalla el resultado final (la multiplicacion entre las sumatorias)
	fmt.Println("Resultado final -->", totalSumRows*totalSumColumns)
}
