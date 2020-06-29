package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("author")
		fmt.Println(f)
		writeFile(args[0], args[1])
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func writeFile(a, b string) {
	// Open file using READ & WRITE permission.
	var file *os.File
	file, err := os.OpenFile("db.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		file, err = os.Create("db.txt")
		if isError(err) {
			return
		}
	}
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = fmt.Fprintln(file, a, b)
	if isError(err) {
		return
	}
	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
