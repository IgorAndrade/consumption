package command

import (
	"fmt"
	"time"

	"github.com/IgorAndrade/consumo_combustivel/api/store"
	"github.com/IgorAndrade/consumo_combustivel/internal/model"
	"github.com/IgorAndrade/consumo_combustivel/internal/service"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "insert new record",
	Example: "add -k 100 -l 14 ",
	Long:    `insert km record`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		f, _ := cmd.Root().PersistentFlags().GetString("file")
		k, _ := cmd.Flags().GetInt64("km")
		l, _ := cmd.Flags().GetFloat64("liters")
		s, _ := cmd.Flags().GetString("station")
		r, _ := cmd.Flags().GetString("route")

		store := store.NewFilestore(f)
		defer store.Close()

		serv := service.NewConsumption(store, store)
		serv.Insert(model.Fuel_Consumption{
			Km:      k,
			Liters:  l,
			Station: s,
			Route:   r,
		})
		fmt.Println("duration", time.Since(start))
	},
}
