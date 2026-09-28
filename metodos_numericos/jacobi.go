package metodos_numericos

import (
	"fmt"
	"metodos-numericos/utils"
)

func metodo_jacobi(matAumentada [][]float64, chute []float64, maxinte int, desvio float64) []float64 {

	chuteant := make([]float64, len(chute))
	copy(chuteant, chute)

	primeiravez := true // Variavel para nao dar problema na primeira vez
	if CriterioLinhas(matAumentada) {
		fmt.Println("Passou no criterio das linhas!")
	} else {
		fmt.Println("Nao passou no criterio das linhas!")
	}

	matA, matB := utils.SepararMatrizAumentada(matAumentada)

	for i := 0; i < maxinte; i++ {
		if !primeiravez && utils.DesvioRelativo(chuteant, chute, desvio) {
			return chute
		}

		copy(chuteant, chute)
		primeiravez = false
		for j := 0; j < len(chute); j++ {
			chute[j] = utils.EncontrarXJacobi(j, matA, chuteant, matB)

		}

	}
	fmt.Println("Numero maximo de iteracoes atingido! Retornando chute")

	return chute

}
