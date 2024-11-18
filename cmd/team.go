package cmd

import "github.com/spf13/cobra"

var teamCmd = &cobra.Command{
	Use:   "team {add | list}",
	Short: "create a new team or list existing ones",
	Long:  "Creates a new team or list existing ones",
}

func init() {
	teamCmd.AddCommand(listTeams)
	teamCmd.AddCommand(addTeam)
}
