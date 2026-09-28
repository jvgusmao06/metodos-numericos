package main

import (
	"fmt"
	"metodos-numericos/metodos_numericos"
)

func main() {

	/*m := [][]float64{
		{3, 2, -5, 1, 8},
		{1, 4, 1, 0, 2},
		{-2, 8, 3, -4, 0},
		{0, -1, 2, 4, 10},
	}

	n := [][]float64{
		{1, 4, -8, -1, 5},
		{-2, 3, -4, -10, -3},
		{-18, 3, 4, -5, 0},
		{1, -6, 0, 3, 10},
	}

	fmt.Println(metodos_numericos.CriterioLinhas(n))
	fmt.Println(metodos_numericos.CriterioLinhas(m))
	*/

	matAumentada := [][]float64{
		{1, 4, -8, -1, 5},
		{-2, 3, -4, -10, -3},
		{-18, 3, 4, -5, 0},
		{1, -6, 0, 3, 10},
	}

	// Chute inicial
	chute := []float64{1, 2, 3, -1}

	// Critério de parada de erro relativo epsilon = 10^-6
	tolerancia := 1e-6

	// Limite máximo de iterações
	maxIteracoes := 1000

	// Executa o método de Jacobi
	resultado := metodos_numericos.GaussJacobi(matAumentada, chute, maxIteracoes, tolerancia)

	// Imprime o resultado
	fmt.Println("Vetor solução:")
	for i, val := range resultado {
		fmt.Printf("x%d = %.6f\n", i+1, val)
	}
}
