package cmd

import "github.com/spf13/cobra"

var projectCmd = &cobra.Command{
	Use:   "project {add | list}",
	Short: "create a new project or list existing ones",
	Long:  "Creates a new project or list existing ones",
}

func init() {
	projectCmd.AddCommand(addProject)
	projectCmd.AddCommand(listProjects)
}
