package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var listTeams = &cobra.Command{
	Use:   "list",
	Short: "list all teams",
	Long:  "Lists all teams in no particular order",
	RunE:  runListTeams,
}

func runListTeams(cmd *cobra.Command, args []string) error {
	repo, err := repository.NewSqlitePngRepository(getDatabaseDsn())
	if err != nil {
		return ErrOpenDatabase{err}
	}

	teams, err := repo.ListTeams()
	if err != nil {
		fmt.Printf("failed to list teams: %s\n", err)
	}
	for _, team := range teams {
		fmt.Println(team)
	}
	return nil
}
