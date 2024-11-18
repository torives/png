package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var listWorkType = &cobra.Command{
	Use:   "list",
	Short: "Lists all work types",
	Run:   runListWorkType,
}

func runListWorkType(cmd *cobra.Command, args []string) {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		fmt.Printf("failed to open database. %s\n", err)
		os.Exit(1)
	}

	workTypes, err := repo.ListWorkTypes()
	if err != nil {
		fmt.Printf("failed to list worktypes. %s\n", err)
	}
	for _, workType := range workTypes {
		fmt.Println(workType)
	}
}
