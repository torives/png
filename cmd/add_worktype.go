package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/model"
	"github.com/torives/png/repository"
)

var addWorkType = &cobra.Command{
	Use:   "add NAME",
	Short: "add a new work type",
	Long: `Adds a new work type to the database. A team name is formed by 
combining any two uppercase letters.

Examples:
MA
ES
IC`,
	RunE: runAddWorkType,
}

func runAddWorkType(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return ErrMissingNameParameter
	}

	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		return ErrOpenDatabase{err}
	}

	name := args[0]
	if err := model.ValidateWorkTypeName(name); err != nil {
		return err
	}

	err = repo.InsertWorkType(model.WorkType{Name: name})
	if err != nil {
		return fmt.Errorf("failed to insert work type: %w\n", err)
	}

	fmt.Printf("Work type \"%s\" was created\n", name)
	return nil
}
