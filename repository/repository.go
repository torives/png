package repository

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	sqliteDriver "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/torives/png/model"
	"modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrations embed.FS

var (
	ErrDuplicatedName = errors.New("duplicated name")

	DB_NAME = "png"
)

type PngRepository interface {
	InsertTeam(team model.Team) error
	GetTeam(name string) (team *model.Team, err error)
	ListTeams() (teams []model.Team, err error)
	InsertWorkType(workType model.WorkType) error
	GetWorkType(name string) (workType *model.WorkType, err error)
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
	err = repo.runMigrations()
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

var (
	insertTeamSql     = `INSERT INTO teams VALUES(NULL, $1)`
	insertWorkTypeSql = `INSERT INTO work_types VALUES(NULL, $1)`
	insertProjectSql  = `INSERT INTO projects VALUES($1, $2, $3)`
	//TODO: find out why parameter substitution does not work with SELECT statements
	selectLastIdSql     = `SELECT id FROM projects WHERE projects.team == $1 AND projects.work_type == $2 ORDER BY id DESC LIMIT 1`
	listAllTeamsSql     = `SELECT name FROM teams`
	listAllWorkTypesSql = `SELECT name FROM work_types`
	listAllProjectsSql  = `SELECT id, team, work_type FROM projects`
	getTeamSql          = `SELECT name FROM teams WHERE teams.name == $1`
	getWorkTypeSql      = `SELECT name FROM work_types WHERE work_types.name == $1`
)

func (r SqlitePngRepository) runMigrations() error {
	driver, err := sqliteDriver.WithInstance(r.db, &sqliteDriver.Config{DatabaseName: DB_NAME})
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", source, DB_NAME, driver)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (r SqlitePngRepository) InsertTeam(team model.Team) error {
	_, err := r.db.Exec(insertTeamSql, team.Name)

	var sqliteErr *sqlite.Error
	if errors.As(err, &sqliteErr) && sqliteErr.Code() == 2067 {
		return ErrDuplicatedName
	} else if err != nil {
		return err
	}

	return nil
}

func (r SqlitePngRepository) GetTeam(name string) (team *model.Team, err error) {
	row := r.db.QueryRow(getTeamSql, name)
	var teamName string

	if err := row.Scan(&teamName); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		return nil, err
	}

	return &model.Team{Name: teamName}, nil
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

	var sqliteErr *sqlite.Error
	if errors.As(err, &sqliteErr) && sqliteErr.Code() == 2067 {
		return ErrDuplicatedName
	} else if err != nil {
		return err
	}

	return nil
}

func (r SqlitePngRepository) GetWorkType(name string) (workType *model.WorkType, err error) {
	row := r.db.QueryRow(getWorkTypeSql, name)
	var workTypeName string

	if err := row.Scan(&workTypeName); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		return nil, err
	}

	return &model.WorkType{Name: workTypeName}, nil
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
	rows, err := r.db.Query(selectLastIdSql, team.Name, workType.Name)
	if err != nil {
		return project, err
	}

	var id int64
	for rows.Next() {
		rows.Scan(&id)
	}
	id += 1

	_, err = r.db.Exec(insertProjectSql, id, team.Name, workType.Name)
	if err != nil {
		return project, err
	}

	p, err := model.NewProject(id, team.Name, workType.Name)
	if err != nil {
		return model.Project{}, err
	}

	return *p, err
}

func (r SqlitePngRepository) ListProjects() (projects []model.Project, err error) {
	rows, err := r.db.Query(listAllProjectsSql)
	if err != nil {
		return nil, err
	}

	var (
		id             int64
		team, workType string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &team, &workType); err != nil {
			return nil, err
		}
		project, err := model.NewProject(id, team, workType)
		if err != nil {
			return nil, err
		}
		projects = append(projects, *project)
	}
	return projects, err
}
