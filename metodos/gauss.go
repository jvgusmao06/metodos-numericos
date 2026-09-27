package metodos_numericos

import "metodos-numericos/utils"

func ResolveSistemaTriangular(matriz [][]float64) []float64 {
	solucao := make([]float64, len(matriz))
	var i int
	var j int
	for i = len(matriz) - 1; i >= 0; i-- {
		temp := matriz[i][len(matriz[i])-1]
		iteracao := len(solucao) - 1
		for j = len(matriz[i]) - 2; j > i; j-- {
			temp -= matriz[i][j] * solucao[iteracao]
			iteracao--
		}
		temp /= matriz[i][j]
		solucao[i] = temp
	}
	return solucao
}

func Gauss(matriz [][]float64) []float64 {
	utils.Escalonamento(matriz)
	return ResolveSistemaTriangular(matriz)
}
