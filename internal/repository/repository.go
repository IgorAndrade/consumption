package repository

import "github.com/IgorAndrade/consumo_combustivel/internal/model"

type Writer interface {
	Insert(model.Fuel_Consumption) error
}

type Reader interface {
	ReadAll() []model.Fuel_Consumption
}
