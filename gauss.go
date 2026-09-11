package main

import (
	"fmt"
	"math"
)

func ImprimeMatriz(matriz [][]float64) {
	for i := range matriz {
		for _, valor := range matriz[i] {
			fmt.Printf("%.4f ", valor)
		}
		fmt.Println()
	}
}

func TrocaLinha(matriz [][]float64, i int, j int) {
	temp_vet := make([]float64, len(matriz[0]))
	copy(temp_vet, matriz[i])
	copy(matriz[i], matriz[j])
	copy(matriz[j], temp_vet)
}

func SubtracaoVet(vet1 []float64, vet2 []float64) {
	for i := 0; i < len(vet1); i++ {
		vet1[i] -= vet2[i]
	}
}

func MultEscalar(vet []float64, e float64) {
	for i := range vet {
		vet[i] *= e
	}
}

func Escalonamento(matriz [][]float64) {
	iteracao := 0
	for i := 0; i < len(matriz)-1; i++ {

		index_maior := iteracao // Grava o indice do futuro pivo

		for j := iteracao; j < len(matriz); j++ { // Pega o pivo e salva indice da linha em que ele estava
			valor_modulo := math.Abs(matriz[j][i])
			if valor_modulo > math.Abs(matriz[index_maior][i]) {
				index_maior = j
			}
		}

		TrocaLinha(matriz, index_maior, iteracao)

		for matador := i + 1; matador < len(matriz); matador++ { // Mata os numeros da coluna. ( Zera eles )
			m := matriz[matador][iteracao] / matriz[iteracao][iteracao]

			tempvet := make([]float64, len(matriz[matador]))
			copy(tempvet, matriz[i])

			MultEscalar(tempvet, m)
			SubtracaoVet(matriz[matador], tempvet)
		}
		iteracao++
	}
}

func ResolveSistemaTriangular(matriz [][]float64) []float64 {
	fmt.Printf("Stay shot")

	solucao := make([]float64, len(matriz))
	var i int
	var j int
	for i = len(matriz) - 1; i >= 0; i-- {
		temp := matriz[i][len(matriz[i])-1]
		iteracao := len(solucao) - 1
		for j = len(matriz[i]) - 2; j > i; j-- {
			temp -= matriz[i][j] * solucao[iteracao]
			fmt.Printf("Tirando %f\n", temp)
			iteracao--
		}
		fmt.Printf("Dividindo %f\n\n", matriz[i][j])
		temp /= matriz[i][j]
		solucao[i] = temp
	}
	return solucao
}

func Gauss(matriz [][]float64) []float64 {
	Escalonamento(matriz)
	return ResolveSistemaTriangular(matriz)
}

func main() {
	m := [][]float64{
		{3, 2, -5, 1, 8},
		{1, 4, 1, 0, 2},
		{-2, 8, 3, -4, 0},
		{0, -1, 2, 4, 10},
	}

	fmt.Println(Gauss(m))
	ImprimeMatriz(m)
}
