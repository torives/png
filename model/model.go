package model

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

var (
	ErrInvalidTeamName     = errors.New("invalid team name")
	ErrInvalidWorkTypeName = errors.New("invalid work type name")
	ErrInvalidProjectId    = errors.New("invalid project id")
	teamNameRegex          = "^[A-Z]{3}$"
	workTypeNameRegex      = "^[A-Z]{2}$"
)

type Team struct {
	Name string
}

func NewTeam(name string) (*Team, error) {
	if err := ValidateTeamName(name); err != nil {
		return nil, err
	}
	return &Team{Name: name}, nil
}

func (t Team) String() string {
	return t.Name
}

func ValidateTeamName(name string) error {
	matched, err := regexp.MatchString(teamNameRegex, name)
	// invalid regex.MatchString invocation, should never happen
	if err != nil {
		log.Fatal(fmt.Errorf("model: failed to validate team name: %w", err))
	}

	if !matched {
		return ErrInvalidTeamName
	}
	return nil
}

type WorkType struct {
	Name string
}

func NewWorkType(name string) (*WorkType, error) {
	if err := ValidateWorkTypeName(name); err != nil {
		return nil, err
	}
	return &WorkType{Name: name}, nil
}

func (wt WorkType) String() string {
	return wt.Name
}

func ValidateWorkTypeName(name string) error {
	matched, err := regexp.MatchString(workTypeNameRegex, name)
	// invalid regex.MatchString invocation, should never happen
	if err != nil {
		log.Fatal(fmt.Errorf("model: failed to validate work type name: %w", err))
	}

	if !matched {
		return ErrInvalidWorkTypeName
	}
	return nil
}

type Project struct {
	Name string
}

func NewProject(id int64, team string, workType string) (*Project, error) {
	if err := ValidateTeamName(team); err != nil {
		return nil, err
	}
	if err := ValidateWorkTypeName(workType); err != nil {
		return nil, err
	}
	if id < 0 {
		return nil, ErrInvalidProjectId
	}
	return &Project{fmt.Sprintf("%s-%s-%d", team, workType, id)}, nil
}

func (p Project) String() string {
	return p.Name
}
