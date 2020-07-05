package service

import (
	"github.com/IgorAndrade/consumo_combustivel/internal/model"
	"github.com/IgorAndrade/consumo_combustivel/internal/repository"
)

type Consumption struct {
	repo repository.Repository
}

func NewConsumption(repo repository.Repository) Consumption {
	return Consumption{repo: repo}
}

func (c Consumption) Insert(fc model.Fuel_Consumption) error {
	list := c.repo.ReadAll()
	fc = calculate(fc, list)
	return c.repo.Insert(fc)
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
	return c.repo.ReadAll()
}
