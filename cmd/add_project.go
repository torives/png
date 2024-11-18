package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torives/png/model"
	"github.com/torives/png/repository"
)

var (
	team                string
	workType            string
	ErrTeamNotFound     = errors.New("team not found")
	ErrWorkTypeNotFound = errors.New("work type not found")
)

var addProject = &cobra.Command{
	Use:   "add {-t | --team} TEAM {-w | --worktype} WORKTYPE",
	Short: "add a new project",
	Long: `Adds a new project. A project has a responsible team and a type of 
work and both are reflected in its name. All project names follow the template 
[AAA-BB-#], where:
  AAA - the code for the team. Check "png team add" command for more info.
  BB - the code for the work type. Check "png worktype add" command for more 
info.
  # - an ever-increasing integer number that uniquely identifies that project.

Examples:
  png project add -t FOR -w PP
  png project add --team FOR --worktype PP`,
	RunE: runAddProject,
}

func runAddProject(cmd *cobra.Command, args []string) error {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		return ErrOpenDatabase{err}
	}

	if err := model.ValidateTeamName(team); err != nil {
		return err
	}

	team, err := repo.GetTeam(team)
	if err != nil {
		return err
	} else if team == nil {
		return ErrTeamNotFound
	}

	if err := model.ValidateWorkTypeName(workType); err != nil {
		return err
	}

	workType, err := repo.GetWorkType(workType)
	if err != nil {
		return err
	} else if workType == nil {
		return ErrWorkTypeNotFound
	}

	project, err := repo.CreateNewProject(*team, *workType)
	if err != nil {
		return fmt.Errorf("failed to create project: %w\n", err)
	}

	fmt.Printf("Project %v was created\n", project)
	return nil
}

func init() {
	addProject.Flags().StringVarP(&team, "team", "t", "", "the team responsible for the project")
	addProject.Flags().StringVarP(&workType, "worktype", "w", "", "the project's type of work")
}
