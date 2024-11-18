package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addTeam = &cobra.Command{
	Use:   "add NAME",
	Short: "Creates a new team with the specified name",
	//TODO: implement add team
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}
