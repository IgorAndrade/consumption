package service

import (
	"fmt"

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
	first := list[:1]
	last := list[len(list)-1:]
	fmt.Println(first, last)
	return c.repo.Insert(fc)
}

func (c Consumption) ReadAll() []model.Fuel_Consumption {
	return c.repo.ReadAll()
}
