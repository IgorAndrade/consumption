package service

import (
	"github.com/IgorAndrade/consumo_combustivel/internal/model"
	"github.com/IgorAndrade/consumo_combustivel/internal/repository"
)

type Consumption struct {
	reader repository.Reader
	writer repository.Writer
}

func NewConsumption(reader repository.Reader, writer repository.Writer) Consumption {
	return Consumption{reader: reader, writer: writer}
}

func (c Consumption) Insert(fc model.Fuel_Consumption) error {
	list := c.reader.ReadAll()
	fc = calculate(fc, list)
	return c.writer.Insert(fc)
}

func calculate(recived model.Fuel_Consumption, list []model.Fuel_Consumption) model.Fuel_Consumption {
	if len(list) == 0 {
		return recived
	}

	first := list[0]
	last := list[len(list)-1]

	recived.KmInterval = recived.Km - last.Km
	recived.KmStart = recived.Km - first.Km
	recived.LitersTotal = totalLiters(append(list, recived))
	recived.Avg = float64(recived.KmInterval) / recived.Liters
	recived.AvgTotal = float64(recived.KmStart) / recived.LitersTotal

	return recived
}

func totalLiters(list []model.Fuel_Consumption) (total float64) {
	for _, l := range list {
		total += l.Liters
	}
	return
}

func (c Consumption) ReadAll() []model.Fuel_Consumption {
	return c.reader.ReadAll()
}
