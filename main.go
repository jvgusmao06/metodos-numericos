package main

import "fmt"
import "metodos-numericos/metodos_numericos"
import "metodos-numericos/utils"

func main() {
	m := [][]float64{
		{3, 2, -5, 1, 8},
		{1, 4, 1, 0, 2},
		{-2, 8, 3, -4, 0},
		{0, -1, 2, 4, 10},
	}

	fmt.Println(metodos_numericos.Gauss(m))
	utils.ImprimeMatriz(m)
}
