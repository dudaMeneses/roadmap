package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [TASK]",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		if f, err := os.OpenFile(time.Now().Format(time.DateOnly)+fileSufix, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			panic(err)
		} else {
			defer f.Close()

			scanner := bufio.NewScanner(f)

			todoTask := "☑️ " + task + "\n"
			doneTask := "✅ " + task + "\n"

			for scanner.Scan() {
				if scanner.Text() == todoTask || scanner.Text() == doneTask {
					fmt.Printf("Task \"%s\" already exists\n", task)
					return
				}
			}

			if _, err := f.WriteString(todoTask); err != nil {
				panic(err)
			} else {
				fmt.Printf("Added \"%s\" to your task list\n", task)
			}
		}
	},
}
