package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		task:=strings.Join(args," ")
		fmt.Printf("Task \"%s\" added to your task list.",task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

//   https://gophercises.com/exercises/task  14:56
