package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const fileSufix = "-tasks.txt"

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(doneCmd)
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI for managing your todo list",
	Long: `Todo is a CLI for managing your todo list where you can:
				- Add a new task
				- List all tasks
				- Remove a task
				- Mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Todo CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
