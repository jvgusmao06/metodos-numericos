package utils

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

func EncontrarXJacobi(ind_linha int, mat [][]float64, vet_x []float64, vet_resp []float64) float64 {
	resp := vet_resp[ind_linha]

	for j := 0; j < len(mat); j++ {
		if j == ind_linha {
			continue
		}
		resp -= (mat[ind_linha][j] * vet_x[j])
	}
	resp /= mat[ind_linha][ind_linha]

	return resp

}
