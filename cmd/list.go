package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var list = &cobra.Command{
	Use:   "list {teams | worktypes}",
	Short: "list database contents",
	Long: `The list command can be used to list all teams and work types present
in the system`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"teams", "worktypes"},
	Run:       runList,
}

func runList(cmd *cobra.Command, args []string) {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		fmt.Printf("failed to open database. %s\n", err)
		os.Exit(1)
	}

	switch args[0] {
	case "teams":
		teams, err := repo.ListTeams()
		if err != nil {
			fmt.Printf("failed to list teams. %s\n", err)
		}
		for _, team := range teams {
			fmt.Println(team)
		}
	case "worktypes":
		workTypes, err := repo.ListWorkTypes()
		if err != nil {
			fmt.Printf("failed to list worktypes. %s\n", err)
		}
		for _, workType := range workTypes {
			fmt.Println(workType)
		}
	}
}

func init() {
	rootCmd.AddCommand(list)
}
