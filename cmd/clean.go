package command

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove file",
	Long:  `remove file and erase all data`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		f, _ := cmd.Root().PersistentFlags().GetString("file")
		if err := os.Remove(f); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("File %s removed successfully. \n", f)
		fmt.Println("duration", time.Since(start))
	},
}
