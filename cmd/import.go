package command

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IgorAndrade/consumo_combustivel/api/store"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import records from excel",
	Long:  `import records from excel`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		//f, _ := cmd.Root().PersistentFlags().GetString("file")
		e, _ := cmd.PersistentFlags().GetString("excel")

		if !strings.HasSuffix(e, "xls") && !strings.HasSuffix(e, "xlsx") {
			log.Fatalln("xls or xlsx required", e)
		}

		store := store.NewExcel(e)
		defer store.Close()
		store.ReadAll()
		// serv := service.NewConsumption(store, store)
		// list := serv.ReadAll()
		// for i, v := range list {
		// 	fmt.Println(i+1, "->", v)
		// }
		fmt.Println("duration", time.Since(start))
	},
}
