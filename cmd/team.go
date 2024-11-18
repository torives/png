package cmd

import "github.com/spf13/cobra"

var teamCmd = &cobra.Command{
	Use:   "team {add | list}",
	Short: "Creates a new team or lists existing ones",
}

func init() {
	teamCmd.AddCommand(listTeams)
	teamCmd.AddCommand(addTeam)
}
