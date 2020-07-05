package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "consumption",
	Short: "fuel consumption short",
	Long:  `fuel consumption long`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	addCmd.Flags().Int64P("km", "k", 0, "kilometers traveled")
	addCmd.Flags().Float64P("liters", "l", 1, "liters consumed")
	addCmd.Flags().StringP("station", "s", "n/a", "liters consumed")
	addCmd.Flags().StringP("route", "r", "n/a", "liters consumed")

	importCmd.Flags().StringP("excel", "e", "excel.xlsx", "excel name")

	rootCmd.PersistentFlags().StringP("file", "f", "db.txt", "store all records")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(importCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
