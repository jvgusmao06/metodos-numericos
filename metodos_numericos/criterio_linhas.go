package metodos_numericos

import (
	"math"
	"metodos-numericos/utils"
)

func alpha(lin []float64, ind_x int) float64 {
	var sum float64

	for i := 0; i < len(lin)-1; i++ {
		if i != ind_x {
			sum += math.Abs(lin[i])
		}
	}

	sum /= math.Abs(lin[ind_x])
	return sum
}

func ChecaLinhas(mat [][]float64) bool {
	for i := range mat {
		a := alpha(mat[i], i)
		if a > 1 {
			return false
		}
	}
	return true
}

func CriterioLinhas(mat [][]float64) bool {
	if ChecaLinhas(mat) {
		return true
	}

	for i := 0; i < len(mat)-1; i++ {
		ind := EncontrarPivo(mat, i)
		if i != ind {
			utils.TrocaLinha(mat, i, ind)
		}
	}

	return ChecaLinhas(mat)
}

func EncontrarPivo(mat [][]float64, ind_linha int) int {
	maior_ind := ind_linha
	maior := math.Abs(mat[ind_linha][ind_linha])

	for i := ind_linha; i < len(mat); i++ {

		if math.Abs(mat[i][ind_linha]) > math.Abs(maior) {
			maior = mat[i][ind_linha]
			maior_ind = i
		}

	}
	return maior_ind
}
