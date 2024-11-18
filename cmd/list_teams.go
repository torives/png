package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var listTeams = &cobra.Command{
	Use:   "list",
	Short: "Lists all teams",
	Run:   runListTeams,
}

func runListTeams(cmd *cobra.Command, args []string) {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		fmt.Printf("failed to open database. %s\n", err)
		os.Exit(1)
	}

	teams, err := repo.ListTeams()
	if err != nil {
		fmt.Printf("failed to list teams. %s\n", err)
	}
	for _, team := range teams {
		fmt.Println(team)
	}
}
