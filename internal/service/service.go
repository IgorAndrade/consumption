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
	return c.repo.Insert(fc)
}

func (c Consumption) ReadAll() []model.Fuel_Consumption {
	return c.repo.ReadAll()
}
