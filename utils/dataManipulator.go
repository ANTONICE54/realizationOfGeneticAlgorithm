package utils

import "gonum.org/v1/gonum/stat/distuv"

func CrossOver(x1, y1, x2, y2 float64, min, max float64) (x, y float64) {
	d := 0.5
	gen := distuv.Uniform{
		Min: -3,
		Max: 1 + d,
	}

	beta := gen.Rand()

	x = beta*x1 + (1-beta)*x2
	y = beta*y1 + (1-beta)*y2

	for y > max || y < min || x > max || x < min {
		beta = gen.Rand()
		x = beta*x1 + (1-beta)*x2
		y = beta*y1 + (1-beta)*y2
	}

	return

}

func Mutation(x1, y1 float64, min, max float64) (x, y float64) {
	gen := distuv.Normal{
		Mu:    0,
		Sigma: (max - min) / 6.,
	}

	mutationCoef := gen.Rand()

	x = x1 + mutationCoef
	y = y1 + mutationCoef
	for x > max || x < min || y > max || y < min {
		mutationCoef = gen.Rand()
		x = x1 + mutationCoef
		y = y1 + mutationCoef
	}

	return

}
