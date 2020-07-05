package store

import (
	"io"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/IgorAndrade/consumo_combustivel/internal/model"
)

type Excel struct {
	file  io.ReadCloser
	start int
}

func NewExcel(name string, start int) *Excel {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		log.Fatalln(err)
	}
	return &Excel{file: file, start: start}
}

func (e Excel) ReadAll() []model.Fuel_Consumption {
	excel, err := excelize.OpenReader(e.file)
	if err != nil {
		log.Fatalln(err)
	}

	sheet := excel.GetSheetName(1)
	records := make([]model.Fuel_Consumption, 0)
	for _, row := range excel.GetRows(sheet)[e.start:] {
		record := make([]string, 0)
		for _, colCell := range row {
			record = append(record, colCell)
		}
		model := convertListToModel(record)
		records = append(records, model)
	}
	return records
}

func (f Excel) Close() error {
	return f.file.Close()
}
