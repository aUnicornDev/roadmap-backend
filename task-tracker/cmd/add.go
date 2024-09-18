package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"dhdhd"},
	Short:   "dsf",
	Long:    "dd",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := AddTask(args[0])
		fmt.Printf(`Task added successfully (ID: %v)`, taskID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
