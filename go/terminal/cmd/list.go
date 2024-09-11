package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if f, err := os.Open(time.Now().Format(time.DateOnly) + fileSufix); err != nil {
			fmt.Println("No tasks for today")
		} else {
			defer f.Close()
			if _, err := io.Copy(os.Stdout, f); err != nil {
				panic(err)
			}
		}
	},
}
