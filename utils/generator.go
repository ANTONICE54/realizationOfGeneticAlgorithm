package utils

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

func RandomGen(min, max float64, n int) [][]float64 {

	gen := distuv.Uniform{
		Min: min,
		Max: max,
	}

	res := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		res[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		res[0][i] = gen.Rand()
		res[1][i] = gen.Rand()
	}
	return res

}

func reverseMap(a map[float64]int) map[int]float64 {
	res := make(map[int]float64)

	for k, v := range a {
		res[v] = k

	}
	return res

}

func HoltonSequence(min, max float64, n int) [][]float64 {
	xMap := make(map[float64]int)
	yMap := make(map[float64]int)

	xCount := 0
	yCount := 0
	diff := max - min
	res := make([][]float64, 2)

	for i := 0; i < 2; i++ {
		res[i] = make([]float64, n)
	}

	i := 1.

	for xCount < n {
		toAdd := 0.
		d := diff / math.Pow(2., i)
		toAdd = min + d
		for toAdd < max && xCount < n {
			if _, ok := xMap[toAdd]; !ok {
				xMap[toAdd] = xCount

				xCount++

			}
			toAdd += d
		}

		i++
	}

	i = 1.

	for yCount < n {
		toAdd := 0.
		d := diff / math.Pow(5., i)
		toAdd = min + d
		for toAdd < max && yCount < n {
			if _, ok := yMap[toAdd]; !ok {
				yMap[toAdd] = yCount

				yCount++

			}
			toAdd += d
		}

		i++
	}

	xMapRev := reverseMap(xMap)
	yMapRev := reverseMap(yMap)

	for j := 0; j < n; j++ {
		res[0][j] = xMapRev[j]
		res[1][j] = yMapRev[j]

	}

	return res

}
