package utils

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
