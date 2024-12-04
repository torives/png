package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/repository"
)

var listProjects = &cobra.Command{
	Use:   "list",
	Short: "list all projects",
	Long:  "Lists all projects in no particular order",
	RunE:  runListProjects,
}

func runListProjects(cmd *cobra.Command, args []string) error {
	repo, err := repository.NewSqlitePngRepository(getDatabaseDsn())
	if err != nil {
		return ErrOpenDatabase{err}
	}

	projects, err := repo.ListProjects()
	if err != nil {
		return fmt.Errorf("failed to list projects: %s\n", err)
	}

	for _, project := range projects {
		fmt.Println(project)
	}
	return nil
}
