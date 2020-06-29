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
	addCmd.Flags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.AddCommand(addCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
