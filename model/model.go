package model

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

var (
	ErrInvalidTeam         = errors.New("invalid team name")
	ErrInvalidWorkTypeName = errors.New("invalid work type name")
	teamNameRegex          = "^[A-Z]{3}$"
	workTypeNameRegex      = "^[A-Z]{2}$"
)

type Team struct {
	Name string
}

func NewTeam(name string) (Team, error) {
	if isValidTeamName(name) {
		return Team{name}, nil
	}
	return Team{}, ErrInvalidTeam
}

func (t Team) String() string {
	return t.Name
}

func isValidTeamName(name string) bool {
	match, err := regexp.MatchString(teamNameRegex, name)
	if err != nil {
		fmt.Printf("model: failed to validate team name. %s", err)
		os.Exit(1)
	}
	return match
}

type WorkType struct {
	Name string
}

func NewWorkType(name string) (WorkType, error) {
	if isValidWorkTypeName(name) {
		return WorkType{name}, nil
	}
	return WorkType{}, ErrInvalidWorkTypeName
}

func (wt WorkType) String() string {
	return wt.Name
}

func isValidWorkTypeName(name string) bool {
	match, err := regexp.MatchString(workTypeNameRegex, name)
	if err != nil {
		fmt.Printf("model: failed to validate workType name: %s", err)
		os.Exit(1)
	}
	return match
}

type Project struct {
	Name string
}

func NewProject(id int64, team string, workType string) Project {
	return Project{fmt.Sprintf("%s-%s-%d", team, workType, id)}
}

func (p Project) String() string {
	return p.Name
}
