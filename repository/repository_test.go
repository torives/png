package repository

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/torives/png/model"
)

// TODO: test with actual database files
var testDsn = "file::memory:?&_pragma=foreign_keys(1)"

func SetupError(err error) error {
	return fmt.Errorf("failed to set up test: %w", err)
}

func TestRepository(t *testing.T) {

	t.Run("ItCreatesARepository", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}
		repo.db.Close()
	})

	t.Run("ItInsertsDefaultDataUponCreation", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}
		defer repo.db.Close()

		teams, err := repo.ListTeams()
		if err != nil {
			t.Fatal(err)
		}

		expectedTeamCount := 4
		if len(teams) != expectedTeamCount {
			t.Fatalf("expected %d teams, got %d", expectedTeamCount, len(teams))
		}

		workTypes, err := repo.ListWorkTypes()
		if err != nil {
			t.Fatal(err)
		}

		expectedWorkTypeCount := 5
		if len(workTypes) != expectedWorkTypeCount {
			t.Fatalf("expected %d teams, got %d", expectedWorkTypeCount, len(teams))
		}
	})

	t.Run("itInsertsANewTeam", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		teams, err := repo.ListTeams()
		if err != nil {
			t.Fatal(err)
		}
		previousTeamCount := len(teams)

		team := model.Team{Name: "TRT"}
		err = repo.InsertTeam(team)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		currentTeams, err := repo.ListTeams()
		if err != nil {
			t.Fatal(err)
		}

		expectedTeamCount := previousTeamCount + 1
		if len(currentTeams) != expectedTeamCount {
			t.Fatalf("expected %d teams, got: %d", expectedTeamCount, len(currentTeams))
		}
	})

	t.Run("itFailsToInsertADuplicateTeam", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		team := model.Team{Name: "FOR"}
		err = repo.InsertTeam(team)
		if err == nil {
			t.Fatalf("unexpected success. %s", err)
		}
	})

	t.Run("itInsertsANewWorkType", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		workTypes, err := repo.ListWorkTypes()
		if err != nil {
			t.Fatal(err)
		}
		previousWorkTypeCount := len(workTypes)

		workType := model.WorkType{Name: "ZZ"}
		err = repo.InsertWorkType(workType)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		currentWorkTypes, err := repo.ListWorkTypes()
		if err != nil {
			t.Fatal(err)
		}

		expectedWorkTypeCount := previousWorkTypeCount + 1
		if len(currentWorkTypes) != expectedWorkTypeCount {
			t.Fatalf("expected %d work types, got: %d", expectedWorkTypeCount, len(currentWorkTypes))
		}
	})

	t.Run("itFailsToInsertADuplicateWorkType", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		workType := model.WorkType{Name: "MA"}
		err = repo.InsertWorkType(workType)
		if err == nil {
			t.Fatalf("unexpected success. %s", err)
		}
	})

	t.Run("itCreatesANewProject", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		project, err := repo.CreateNewProject(model.Team{Name: "FOR"}, model.WorkType{Name: "MA"})
		if err != nil {
			t.Fatalf("unexpected failure. %s", err)
		}

		expectedName := "FOR-MA-1"
		if project.Name != expectedName {
			t.Fatalf("expected project name to be %s, got %s", expectedName, project.Name)
		}
	})

	t.Run("itCreatesANewProjectWhenTeamAndWorkTypeAreTheSame", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		team := model.Team{Name: "FOR"}
		workType := model.WorkType{Name: "MA"}
		project1, err := repo.CreateNewProject(team, workType)
		if err != nil {
			t.Fatalf("unexpected failure. %s", err)
		}

		project2, err := repo.CreateNewProject(team, workType)
		if err != nil {
			t.Fatalf("unexpected failure. %s", err)
		}

		if project1.Name == project2.Name {
			t.Fatalf("expected project names to be different. %s", project1.Name)
		}
	})

	t.Run("itIncrementsProjectIdByOneForTheSameTeamAndWorkType", func(t *testing.T) {
		repo, err := NewSqlitePngRepository(testDsn)
		if err != nil {
			t.Fatal(err)
		}

		team := model.Team{Name: "FOR"}
		workType := model.WorkType{Name: "MA"}
		project1, err := repo.CreateNewProject(team, workType)
		if err != nil {
			t.Fatalf("unexpected failure. %s", err)
		}

		project2, err := repo.CreateNewProject(team, workType)
		if err != nil {
			t.Fatalf("unexpected failure. %s", err)
		}

		project1Id, err := idFromProjectName(project1.Name)
		if err != nil {
			t.Fatal(err)
		}
		project2Id, err := idFromProjectName(project2.Name)
		if err != nil {
			t.Fatal(err)
		}

		if project2Id != project1Id+1 {
			t.Fatalf("expected second id to be %d, but got %d", project1Id+1, project2Id)
		}
	})
}

func idFromProjectName(name string) (int64, error) {
	projectIdStr := name[len(name)-1:]
	projectId, err := strconv.ParseInt(projectIdStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return projectId, nil
}
