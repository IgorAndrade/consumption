package service

import (
	"reflect"
	"testing"

	"github.com/IgorAndrade/consumo_combustivel/internal/model"
)

func Test_calculate(t *testing.T) {
	type args struct {
		recived model.Fuel_Consumption
		list    []model.Fuel_Consumption
	}
	tests := []struct {
		name string
		args args
		want model.Fuel_Consumption
	}{
		{
			name: "calc first line",
			args: args{
				recived: model.Fuel_Consumption{
					Km:      119266,
					Route:   "Helvetia",
					Station: "Londres",
				},
				list: make([]model.Fuel_Consumption, 0),
			},
			want: model.Fuel_Consumption{
				Km:      119266,
				Route:   "Helvetia",
				Station: "Londres",
			},
		},
		{
			name: "calc second line",
			args: args{
				recived: model.Fuel_Consumption{
					Km:      119713,
					Liters:  37.09,
					Route:   "Helvetia",
					Station: "Londres",
				},
				list: []model.Fuel_Consumption{
					{
						Km:      119266,
						Route:   "Helvetia",
						Station: "Londres",
					},
				},
			},
			want: model.Fuel_Consumption{
				Km:          119713,
				KmInterval:  447,
				KmStart:     447,
				Liters:      37.09,
				LitersTotal: 37.09,
				Avg:         12.05176597465624,
				AvgTotal:    12.05176597465624,
				Route:       "Helvetia",
				Station:     "Londres",
			},
		},
		{
			name: "calc 3 line",
			args: args{
				recived: model.Fuel_Consumption{
					Km:      120188,
					Liters:  37.403,
					Route:   "Helvetia",
					Station: "Londres",
				},
				list: []model.Fuel_Consumption{
					{
						Km:      119266,
						Route:   "Helvetia",
						Station: "Londres",
					},
					{
						Km:          119713,
						KmInterval:  447,
						KmStart:     447,
						Liters:      37.09,
						LitersTotal: 37.09,
						Avg:         12.05176597465624,
						AvgTotal:    12.05176597465624,
						Route:       "Helvetia",
						Station:     "Londres",
					},
				},
			},
			want: model.Fuel_Consumption{
				Km:          120188,
				KmInterval:  475,
				KmStart:     922,
				Liters:      37.403,
				LitersTotal: 74.493,
				Avg:         12.699516081597734,
				AvgTotal:    12.377001865947136,
				Route:       "Helvetia",
				Station:     "Londres",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.recived, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
