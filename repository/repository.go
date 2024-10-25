package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/torives/png/model"
	"modernc.org/sqlite"
)

type PngRepository interface {
	InsertTeam(team model.Team) error
	ListTeams() (teams []model.Team, err error)
	InsertWorkType(workType model.WorkType) error
	ListWorkTypes() (workTypes []model.WorkType, err error)
	CreateNewProject(team model.Team, workType model.WorkType) (project model.Project, err error)
}

type SqlitePngRepository struct {
	db *sql.DB
}

// TODO: maybe another New func receiving a *sql.DB that doesn't insert data?
func NewSqlitePngRepository(dsn string) (*SqlitePngRepository, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	repo := SqlitePngRepository{db}
	err = repo.insertInitialData()
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

var (
	createTeamsTableSql      = `CREATE TABLE teams(id INTEGER PRIMARY KEY NOT NULL, name VARCHAR(3) UNIQUE)`
	createWorkTypesTableSql  = `CREATE TABLE work_types(id INTEGER PRIMARY KEY NOT NULL, name VARCHAR(2) UNIQUE)`
	createProjectsTableSql   = `CREATE TABLE projects(id INTEGER PRIMARY KEY AUTOINCREMENT, team TEXT NOT NULL, work_type TEXT NOT NULL, FOREIGN KEY(team) REFERENCES teams(name) FOREIGN KEY(work_type) REFERENCES work_types(name))`
	createTeamFkIndexSql     = `CREATE INDEX team_index ON projects(team)`
	createWorkTypeFkIndexSql = `CREATE INDEX work_type_index ON projects(team)`
	insertTeamSql            = `INSERT INTO teams VALUES(NULL, $1)`
	insertWorkTypeSql        = `INSERT INTO work_types VALUES(NULL, $1)`
	insertProjectSql         = `INSERT INTO projects VALUES(NULL, $1, $2) RETURNING id`
	//TODO: find out why parameter substitution does not work with SELECT statements
	listAllTeamsSql     = `SELECT name FROM teams`
	listAllWorkTypesSql = `SELECT name FROM work_types`
)

// FIXME: add migrations file/tooling
func (r SqlitePngRepository) insertInitialData() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var sqlErr *sqlite.Error

	//create tables
	_, err = tx.Exec(createTeamsTableSql)
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 1 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(createWorkTypesTableSql)
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 1 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(createProjectsTableSql)
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 1 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}

	// create indexes
	_, err = tx.Exec(createTeamFkIndexSql)
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 1 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(createWorkTypeFkIndexSql)
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 1 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}

	// populate teams table
	_, err = tx.Exec(insertTeamSql, "FOR")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertTeamSql, "ANA")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertTeamSql, "MIC")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertTeamSql, "PRO")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}

	// populate work_type table
	_, err = tx.Exec(insertWorkTypeSql, "MA")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertWorkTypeSql, "ES")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertWorkTypeSql, "IC")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertWorkTypeSql, "PT")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}
	_, err = tx.Exec(insertWorkTypeSql, "PP")
	if err != nil {
		if !errors.As(err, &sqlErr) || sqlErr.Code() != 2067 {
			return errors.Join(
				fmt.Errorf("tx exec: %w", err),
				tx.Rollback(),
			)
		}
	}

	return tx.Commit()
}

func (r SqlitePngRepository) InsertTeam(team model.Team) error {
	_, err := r.db.Exec(insertTeamSql, team.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r SqlitePngRepository) ListTeams() (teams []model.Team, err error) {
	rows, err := r.db.Query(listAllTeamsSql)
	if err != nil {
		return nil, err
	}

	var team model.Team
	for rows.Next() {
		if err = rows.Scan(&team.Name); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, err
}

func (r SqlitePngRepository) InsertWorkType(workType model.WorkType) error {
	_, err := r.db.Exec(insertWorkTypeSql, workType.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r SqlitePngRepository) ListWorkTypes() (workTypes []model.WorkType, err error) {
	rows, err := r.db.Query(listAllWorkTypesSql)
	if err != nil {
		return nil, err
	}

	var workType model.WorkType
	for rows.Next() {
		if err = rows.Scan(&workType.Name); err != nil {
			return nil, err
		}
		workTypes = append(workTypes, workType)
	}
	return workTypes, err
}

func (r SqlitePngRepository) CreateNewProject(
	team model.Team,
	workType model.WorkType,
) (project model.Project, err error) {
	result, err := r.db.Exec(insertProjectSql, team.Name, workType.Name)
	if err != nil {
		return project, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return project, err
	}

	return model.NewProject(id, team.Name, workType.Name), nil
}
