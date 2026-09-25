package utils

import "math"

func MultEscalar(vet []float64, e float64) {
	for i := range vet {
		vet[i] *= e
	}
}

func SomaVet(vet1 []float64, vet2 []float64) {
	for i := 0; i < len(vet1); i++ {
		vet1[i] += vet2[i]
	}
}

func SubtracaoVet(vet1 []float64, vet2 []float64) {
	for i := 0; i < len(vet1); i++ {
		vet1[i] -= vet2[i]
	}
}

func NormaInfinito(vet []float64) float64 {
	maior := math.Abs(vet[0])
	for i := 1; i < len(vet); i++ {
		val := math.Abs(vet[i])
		if val > maior {
			maior = val
		}
	}
	return maior
}

func DesvioRelativo(vetant []float64, vetatual []float64, tolerancia float64) bool {
	vetdiferenca := make([]float64, len(vetant))

	copy(vetdiferenca, vetatual)

	SubtracaoVet(vetdiferenca, vetant)
	norma := NormaInfinito(vetdiferenca)
	normaatual := NormaInfinito(vetatual)
	if normaatual == 0 {
		return norma < tolerancia
	}

	resp := norma / normaatual

	if resp < tolerancia {
		return true
	}
	return false

}
