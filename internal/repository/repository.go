package repository

import "github.com/IgorAndrade/consumo_combustivel/internal/model"

type Repository interface {
	Insert(model.Fuel_Consumption) error
	ReadAll() []model.Fuel_Consumption
}
