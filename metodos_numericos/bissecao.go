package metodos_numericos

import "math"

func Bissecao(intervalo [2]float64, f func(float64) float64, e float64) [2]float64 {
	a := intervalo[0]
	b := intervalo[1]

	m := (a + b) / 2

	fa := f(a)
	fm := f(m)

	if fm == 0 {
		return [2]float64{m, m}
	}

	if fa*fm < 0 {
		intervalo[1] = m
	} else {
		intervalo[0] = m
	}

	if math.Abs(intervalo[1]-intervalo[0]) < e {
		return intervalo
	}

	return Bissecao(intervalo, f, e)
}
