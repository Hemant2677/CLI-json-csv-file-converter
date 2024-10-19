package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// parseCmd represents the base command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A CLI tool to display and convert data from CSV or JSON files",
	Long:  "This CLI reads and prints data from CSV or JSON files using the specified subcommands and can convert between formats, saving the output in new files.",
}

func Execute() {
	if err := parseCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
