package store

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/IgorAndrade/consumo_combustivel/internal/model"
)

type Filestore struct {
	file *os.File
}

func NewFilestore(name string) *Filestore {
	var file *os.File
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if os.IsNotExist(err) {
		file, err = os.Create(name)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return &Filestore{file: file}
}

func (f Filestore) Insert(fc model.Fuel_Consumption) error {
	_, err := fmt.Fprintf(f.file, "%v;%v;%v;%v;%v;%v;%v;%v;%v\n",
		fc.Km,
		fc.KmInterval,
		fc.KmStart,
		fc.Liters,
		fc.LitersTotal,
		fc.Avg,
		fc.AvgTotal,
		fc.Station,
		fc.Route)
	if err != nil {
		return err
	}
	// Save file changes.
	if err = f.file.Sync(); err != nil {
		return err
	}
	fmt.Println("File Updated Successfully.", fc)
	return nil
}

func (f Filestore) ReadAll() []model.Fuel_Consumption {
	list := make([]model.Fuel_Consumption, 0)

	scanner := bufio.NewScanner(f.file)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), ";")
		list = append(list, convertListToModel(spl))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func convertListToModel(s []string) model.Fuel_Consumption {
	km, _ := strconv.ParseInt(s[0], 10, 64)
	kmI, _ := strconv.ParseInt(s[1], 10, 64)
	kmS, _ := strconv.ParseInt(s[2], 10, 64)
	ltl, _ := strconv.ParseFloat(s[3], 64)
	ltlT, _ := strconv.ParseFloat(s[4], 64)
	avg, _ := strconv.ParseFloat(s[5], 64)
	avgT, _ := strconv.ParseFloat(s[6], 64)
	return model.Fuel_Consumption{
		Km:          km,
		KmInterval:  kmI,
		KmStart:     kmS,
		Liters:      ltl,
		LitersTotal: ltlT,
		Avg:         avg,
		AvgTotal:    avgT,
		Station:     s[7],
		Route:       s[8],
	}
}

func (f Filestore) Close() error {
	return f.file.Close()
}
