package command

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IgorAndrade/consumo_combustivel/api/store"
	"github.com/IgorAndrade/consumo_combustivel/internal/service"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import records from excel",
	Long:  `import records from excel`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		f, _ := cmd.Root().PersistentFlags().GetString("file")
		e, _ := cmd.Flags().GetString("excel")
		l, _ := cmd.Flags().GetInt("line")

		if !strings.HasSuffix(e, "xls") && !strings.HasSuffix(e, "xlsx") {
			log.Fatalln("xls or xlsx required", e)
		}

		file := store.NewFilestore(f)
		excel := store.NewExcel(e, l)
		defer file.Close()
		defer excel.Close()

		serv := service.NewConsumption(excel, file)
		if err := serv.Import(); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("duration", time.Since(start))
	},
}
