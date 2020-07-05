package store

import (
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/IgorAndrade/consumo_combustivel/internal/model"
)

type Excel struct {
	file *os.File
}

func NewExcel(name string) *Excel {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		log.Fatalln(err)
	}
	return &Excel{file: file}
}

func (e Excel) ReadAll() []model.Fuel_Consumption {
	excel, err := excelize.OpenReader(e.file)
	if err != nil {
		log.Fatalln(err)
	}

	sheet := excel.GetSheetName(0)
	for _, row := range excel.GetRows(sheet) {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
	return nil
}

func (f Excel) Close() error {
	return f.file.Close()
}
