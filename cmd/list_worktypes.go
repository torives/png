package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var listWorkTypes = &cobra.Command{
	Use:   "list",
	Short: "list all work types",
	Long:  "Lists all work types in no particular order",
	RunE:  runListWorkTypes,
}

func runListWorkTypes(cmd *cobra.Command, args []string) error {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		return ErrOpenDatabase{err}
	}

	workTypes, err := repo.ListWorkTypes()
	if err != nil {
		fmt.Printf("failed to list worktypes: %s\n", err)
	}

	for _, workType := range workTypes {
		fmt.Println(workType)
	}
	return nil
}
