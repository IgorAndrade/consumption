package command

import (
	"fmt"
	"time"

	"github.com/IgorAndrade/consumo_combustivel/api/store"
	"github.com/IgorAndrade/consumo_combustivel/internal/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all records",
	Example: "list",
	Long:    `list all records`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		f, _ := cmd.Root().PersistentFlags().GetString("file")

		store := store.NewFilestore(f)
		defer store.Close()
		serv := service.NewConsumption(store, store)
		list := serv.ReadAll()
		for i, v := range list {
			fmt.Println(i+1, "->", v)
		}
		fmt.Println("duration", time.Since(start))
	},
}
