package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"laba1/functions"
	"laba1/utils"
)

func main() {

	file := excelize.NewFile()
	if err := file.SaveAs("results.xlsx"); err != nil {
		fmt.Println(err)
	}

	functions.Mihalevich.GeneratePopulation = utils.RandomGen

	res, _, _, _ := functions.Mihalevich.AlgorithmExecution(100, 10)
	WriteToFile(res, "Mihalevich10RandHen")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 20)
	WriteToFile(res, "Mihalevich20RandHen")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 40)
	WriteToFile(res, "Mihalevich40RandHen")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 50)
	WriteToFile(res, "Mihalevich50RandHen")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 100)
	WriteToFile(res, "Mihalevich100RandHen")

	functions.Mihalevich.GeneratePopulation = utils.HoltonSequence

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 10)
	WriteToFile(res, "Mihalevich10Holton")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 20)
	WriteToFile(res, "Mihalevich20Holton")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 40)
	WriteToFile(res, "Mihalevich40Holton")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 50)
	WriteToFile(res, "Mihalevich50Holton")

	res, _, _, _ = functions.Mihalevich.AlgorithmExecution(100, 100)
	WriteToFile(res, "Mihalevich100Holton")

	functions.CrossInTray.GeneratePopulation = utils.RandomGen

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 10)
	WriteToFile(res, "CrossInTray10RandHen")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 20)
	WriteToFile(res, "CrossInTray20RandHen")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 40)
	WriteToFile(res, "CrossInTray40RandHen")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 50)
	WriteToFile(res, "CrossInTray50RandHen")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 100)
	WriteToFile(res, "CrossInTray100RandHen")

	functions.CrossInTray.GeneratePopulation = utils.HoltonSequence

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 10)
	WriteToFile(res, "CrossInTray10Holton")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 20)
	WriteToFile(res, "CrossInTrayh20Holton")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 40)
	WriteToFile(res, "CrossInTray40Holton")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 50)
	WriteToFile(res, "CrossInTray50Holton")

	res, _, _, _ = functions.CrossInTray.AlgorithmExecution(100, 100)
	WriteToFile(res, "CrossInTray100Holton")

}

func WriteToFile(data [][]float64, sheetName string) {
	file, _ := excelize.OpenFile("results.xlsx")

	headers := []string{"epoch", "x", "y", "functionValue"}
	file.NewSheet(sheetName)
	for i, header := range headers {
		file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	file.SetColWidth(sheetName, "B", "D", 30.)
	for j := 0; j < len(data[0]); j++ {
		dataRow := j + 2

		file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune(65)), dataRow), (j+1)*10)
		file.SetCellFloat(sheetName, fmt.Sprintf("%s%d", string(rune(66)), dataRow), data[0][j], 20, 64)
		file.SetCellFloat(sheetName, fmt.Sprintf("%s%d", string(rune(67)), dataRow), data[1][j], 20, 64)
		file.SetCellFloat(sheetName, fmt.Sprintf("%s%d", string(rune(68)), dataRow), data[2][j], 20, 64)
	}

	if err := file.SaveAs("results.xlsx"); err != nil {
		fmt.Println(err)
	}

}
