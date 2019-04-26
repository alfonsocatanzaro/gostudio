package cmd

//go get -u github.com/spf13/cobra/cobra
import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager.",
}
