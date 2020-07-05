package model

import "fmt"

type Fuel_Consumption struct {
	Km          int64
	KmInterval  int64
	KmStart     int64
	Liters      float64
	LitersTotal float64
	Avg         float64
	AvgTotal    float64
	Station     string
	Route       string
}

func (fc Fuel_Consumption) String() string {
	return fmt.Sprintf("Km: %v, Km (Interval): %v, Liters: %v, Liters (Total): %v, AVG %v, AVG (Total): %v",
		fc.Km,
		fc.KmInterval,
		fc.Liters,
		fc.LitersTotal,
		fc.Avg,
		fc.AvgTotal,
	)
}
