package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [TASK]",
	Short: "Mark a task as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		if f, err := os.OpenFile(time.Now().Format(time.DateOnly)+fileSufix, os.O_RDWR, 0666); err != nil {
			fmt.Println("No tasks for today")
		} else {
			defer f.Close()

			todoTask := "☑️ " + task + "\n"
			doneTask := "✅ " + task + "\n"

			scanner := bufio.NewScanner(f)
			var lines []string
			for scanner.Scan() {
				if scanner.Text() == todoTask {
					lines = append(lines, doneTask)
				} else {
					lines = append(lines, scanner.Text())
				}
			}

			if err := f.Truncate(0); err != nil {
				panic(err)
			}

			if _, err := f.Seek(0, 0); err != nil {
				panic(err)
			}

			for _, line := range lines {
				if _, err := f.WriteString(line + "\n"); err != nil {
					panic(err)
				}
			}
		}
	},
}
