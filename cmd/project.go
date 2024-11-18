package cmd

import "github.com/spf13/cobra"

var projectCmd = &cobra.Command{
	Use:   "project {add | list}",
	Short: "Creates a new project or lists existing ones",
}

func init() {
	projectCmd.AddCommand(addProject)
	projectCmd.AddCommand(listProjects)
}
