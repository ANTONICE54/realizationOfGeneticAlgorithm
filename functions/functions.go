package functions

import (
	"laba1/utils"
	"math"
	"math/rand"
)

type functionToUse func(float64, float64) float64
type generationMethod func(float64, float64, int) [][]float64

type functionMetadata struct {
	value float64
	index int
	rang  int
}

type Method struct {
	Min, Max           float64
	ExecuteFunction    functionToUse
	GeneratePopulation generationMethod
}

var CrossInTray Method = Method{
	Min:             -10.,
	Max:             10.,
	ExecuteFunction: CrossInTrayFunction,
}

var Mihalevich Method = Method{
	Min:             0.,
	Max:             math.Pi,
	ExecuteFunction: MihalevichFunction,
}

func MihalevichFunction(x, y float64) float64 {
	return -math.Sin(x)*math.Pow(math.Sin((x*x)/math.Pi), 20) - math.Sin(y)*math.Pow(math.Sin((2*y*y)/math.Pi), 20)

}

func CrossInTrayFunction(x, y float64) float64 {
	return -0.0001 * math.Pow(math.Abs(math.Sin(x)*math.Sin(y)*math.Exp(math.Abs(100.-math.Sqrt(x*x+y*y)/math.Pi)))+1., 0.1)
}

func (mth Method) AlgorithmExecution(epoch int, populationSize int) (res [][]float64, x, y, f float64) {

	res = make([][]float64, 3)

	currentGeneration := utils.HoltonSequence(mth.Min, mth.Max, populationSize)

	var funcValues []functionMetadata

	for i := 0; i < epoch; i++ {
		funcValues = make([]functionMetadata, populationSize)
		for k := 0; k < populationSize; k++ {
			funcValues[k].value = mth.ExecuteFunction(currentGeneration[0][k], currentGeneration[1][k])
			funcValues[k].index = k
		}
		sortFunctionMetadata(funcValues)

		if (i+1)%10 == 0 || i == epoch-1 {
			res[0] = append(res[0], currentGeneration[0][funcValues[0].index])
			res[1] = append(res[1], currentGeneration[1][funcValues[0].index])
			res[2] = append(res[2], funcValues[0].value)
		}

		if i == epoch-1 {
			break
		}

		for k := 0; k < populationSize; k++ {
			funcValues[k].rang = populationSize - k
		}

		bestAndRandAmount := populationSize / 3
		crossOverAmount := populationSize / 3
		mutationAmount := populationSize - (bestAndRandAmount + crossOverAmount)

		bestAndRandPoints := make([][]float64, 2)
		bestAndRandPoints[0] = make([]float64, bestAndRandAmount)
		bestAndRandPoints[1] = make([]float64, bestAndRandAmount)

		for j := 0; j < bestAndRandAmount/5; j++ {
			bestAndRandPoints[0][j] = currentGeneration[0][funcValues[j].index]
			bestAndRandPoints[1][j] = currentGeneration[1][funcValues[j].index]
		}

		for j := bestAndRandAmount / 5; j < bestAndRandAmount; j++ {
			rnd := rand.Int() % (populationSize - bestAndRandAmount/5 - 1)
			bestAndRandPoints[0][j] = currentGeneration[0][funcValues[rnd].index]
			bestAndRandPoints[1][j] = currentGeneration[1][funcValues[rnd].index]
		}

		crossOverPoints := make([][]float64, 2)
		crossOverPoints[0] = make([]float64, crossOverAmount)
		crossOverPoints[1] = make([]float64, crossOverAmount)

		for j := 0; j < crossOverAmount; j++ {
			rnd1 := rand.Int() % (bestAndRandAmount - 1)
			rnd2 := rand.Int() % (bestAndRandAmount - 1)
			for rnd1 == rnd2 {
				rnd2 = rand.Int() % (bestAndRandAmount - 1)
			}
			xRes, yRes := utils.CrossOver(bestAndRandPoints[0][rnd1], bestAndRandPoints[1][rnd1], bestAndRandPoints[0][rnd2], bestAndRandPoints[1][rnd2], mth.Min, mth.Max)

			crossOverPoints[0][j], crossOverPoints[1][j] = xRes, yRes

		}

		mutationPoints := make([][]float64, 2)
		mutationPoints[0] = make([]float64, mutationAmount)
		mutationPoints[1] = make([]float64, mutationAmount)

		for j := 0; j < mutationAmount; j++ {
			rnd := rand.Int() % (bestAndRandAmount - 1)

			xRes, yRes := utils.Mutation(bestAndRandPoints[0][rnd], bestAndRandPoints[1][rnd], mth.Min, mth.Max)

			mutationPoints[0][j], mutationPoints[1][j] = xRes, yRes

		}
		funcValues = nil
		currentGeneration = nil
		currentGeneration = make([][]float64, 2)

		currentGeneration[0] = append(currentGeneration[0], bestAndRandPoints[0]...)
		currentGeneration[0] = append(currentGeneration[0], crossOverPoints[0]...)
		currentGeneration[0] = append(currentGeneration[0], mutationPoints[0]...)

		currentGeneration[1] = append(currentGeneration[1], bestAndRandPoints[1]...)
		currentGeneration[1] = append(currentGeneration[1], crossOverPoints[1]...)
		currentGeneration[1] = append(currentGeneration[1], mutationPoints[1]...)

	}
	f = funcValues[0].value
	x = currentGeneration[0][funcValues[0].index]
	y = currentGeneration[1][funcValues[0].index]
	return
}

func sortFunctionMetadata(funcMD []functionMetadata) {
	for i := 0; i < len(funcMD); i++ {
		for j := i + 1; j < len(funcMD); j++ {
			if funcMD[j].value <= funcMD[i].value {
				funcMD[j], funcMD[i] = funcMD[i], funcMD[j]
			}
		}
	}
}
