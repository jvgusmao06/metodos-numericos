package metodos_numericos

import (
	"fmt"
	"metodos-numericos/utils"
)

func EncontrarX(ind_linha int, mat [][]float64, vet_x []float64, vet_resp []float64) float64 {
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

func GaussJacobi(matAumentada [][]float64, chute []float64, maxinte int, desvio float64) []float64 {

	chuteant := make([]float64, len(chute))
	copy(chuteant, chute)

	primeiravez := true // Variavel para nao dar problema na primeira vez
	if CriterioLinhas(matAumentada) {
		fmt.Println("Passou no criterio das linhas!")
	} else {
		fmt.Println("Nao passou no criterio das linhas!")
	}

	//matA, matB := utils.SepararMatrizAumentada(matAumentada)

	for i := 0; i < maxinte; i++ {
		if !primeiravez && utils.DesvioRelativo(chuteant, chute, desvio) {
			return chute
		}

		copy(chuteant, chute)
		primeiravez = false
		for j := 0; j < len(chute); j++ {
			//chute[j] = EncontrarX(j, matA, chuteant, matB)

		}

	}
	fmt.Println("Numero maximo de iteracoes atingido! Retornando chute")

	return chute

}
