package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/model"
	"github.com/torives/png/repository"
)

var ErrMissingNameParameter = errors.New("missing name parameter")

var addTeam = &cobra.Command{
	Use:   "add NAME",
	Short: "add a new team",
	Long: `Adds a new team to the database. A team name is formed by combining 
any three uppercase letters.

Examples:
FOR
ANA
MIC`,
	RunE: runAddTeam,
}

func runAddTeam(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Usage()
		return ErrMissingNameParameter
	}

	repo, err := repository.NewSqlitePngRepository(getDatabaseDsn())
	if err != nil {
		return ErrOpenDatabase{err}
	}

	name := args[0]
	if err := model.ValidateTeamName(name); err != nil {
		return err
	}

	err = repo.InsertTeam(model.Team{Name: name})
	if err != nil {
		return fmt.Errorf("failed to insert team: %w\n", err)
	}

	fmt.Printf("Team \"%s\" was created\n", name)
	return nil
}
