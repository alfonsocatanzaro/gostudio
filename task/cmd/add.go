package cmd

import (
	"fmt"
	"github.com/gostudio/task/db"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Task \"%s\" added to your task list.", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

//   https://gophercises.com/exercises/task  14:56
