package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listProjects = &cobra.Command{
	Use:   "list",
	Short: "Lists all projects",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: implement list projects
		fmt.Println("TODO")
	},
}
