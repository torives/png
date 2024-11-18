package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/torives/png/model"
	"github.com/torives/png/repository"
)

var (
	team                string
	workType            string
	ErrMissingTeam      = errors.New("missing team parameter")
	ErrTeamNotFound     = errors.New("team not found")
	ErrMissingWorkType  = errors.New("missing worktype parameter")
	ErrWorkTypeNotFound = errors.New("worktype not found")
)

var addProject = &cobra.Command{
	Use:   "add [-t | --team] TEAM [-w | --worktype] WORKTYPE",
	Short: "Creates a new project",
	Run:   runAddProject,
}

func runAddProject(cmd *cobra.Command, args []string) {
	repo, err := repository.NewSqlitePngRepository(databaseDsn)
	if err != nil {
		fmt.Printf("failed to open database. %s\n", err)
		os.Exit(1)
	}

	if err = validateTeam(repo); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = validateWorkType(repo); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	project, err := repo.CreateNewProject(model.Team{Name: team}, model.WorkType{Name: workType})
	if err != nil {
		fmt.Printf("failed to create project. %s\n", err)
		os.Exit(1)
	}

	fmt.Println(project)
}

func validateTeam(repo repository.PngRepository) error {
	if team == "" {
		return ErrMissingTeam
	}

	//TODO: add repo.GetTeam() method
	teams, err := repo.ListTeams()
	if err != nil {
		return err
	}

	for _, storedTeam := range teams {
		if storedTeam.Name == team {
			return nil
		}
	}
	return ErrTeamNotFound
}

func validateWorkType(repo repository.PngRepository) error {
	if workType == "" {
		return ErrMissingWorkType
	}

	//TODO: add repo.GetWorkType() method
	workTypes, err := repo.ListWorkTypes()
	if err != nil {
		return err
	}

	for _, storedWorkType := range workTypes {
		if storedWorkType.Name == workType {
			return nil
		}
	}
	return ErrWorkTypeNotFound
}

func init() {
	addProject.Flags().StringVarP(&team, "team", "t", "", "the team responsible for the project")
	addProject.Flags().StringVarP(&workType, "worktype", "w", "", "the project's type of work")
}
