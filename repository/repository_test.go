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
		repo := newSqliteRepository(t, testDsn)

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
		repo := newSqliteRepository(t, testDsn)

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
		repo := newSqliteRepository(t, testDsn)

		team := model.Team{Name: "FOR"}
		err := repo.InsertTeam(team)
		if err == nil {
			t.Fatalf("unexpected success. %s", err)
		}
	})

	t.Run("itInsertsANewWorkType", func(t *testing.T) {
		repo := newSqliteRepository(t, testDsn)

		workTypes, err := repo.ListWorkTypes()
		if err != nil {
			t.Fatal(err)
		}
		previousWorkTypeCount := len(workTypes)

		workType := model.WorkType{Name: "ZZ"}
		err = repo.InsertWorkType(workType)
		if err != nil {
			t.Fatalf("expected to insert a new worktype. %s", err)
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
		repo := newSqliteRepository(t, testDsn)

		workType := model.WorkType{Name: "MA"}
		err := repo.InsertWorkType(workType)
		if err == nil {
			t.Fatal("expected insertion of duplicate worktype to fail, but it did not")
		}
	})

	t.Run("itCreatesANewProject", func(t *testing.T) {
		repo := newSqliteRepository(t, testDsn)

		project, err := repo.CreateNewProject(model.Team{Name: "FOR"}, model.WorkType{Name: "MA"})
		if err != nil {
			t.Fatalf("expected project creation to succeed. %s", err)
		}

		expectedName := "FOR-MA-1"
		if project.Name != expectedName {
			t.Fatalf("expected project name to be %s, got %s", expectedName, project.Name)
		}
	})

	t.Run("itCreatesANewProjectWhenTeamAndWorkTypeAreTheSame", func(t *testing.T) {
		repo := newSqliteRepository(t, testDsn)

		team := model.Team{Name: "FOR"}
		workType := model.WorkType{Name: "MA"}
		project1 := createNewProject(t, repo, team, workType)
		project2 := createNewProject(t, repo, team, workType)

		if project1.Name == project2.Name {
			t.Fatalf("expected projects %s and %s to have different names", project1.Name, project2.Name)
		}
	})

	t.Run("itIncrementsProjectIdByOneForTheSameTeamAndWorkType", func(t *testing.T) {
		repo := newSqliteRepository(t, testDsn)

		team := model.Team{Name: "FOR"}
		workType := model.WorkType{Name: "MA"}
		project1 := createNewProject(t, repo, team, workType)
		project2 := createNewProject(t, repo, team, workType)

		project1Id := idFromProject(t, project1)
		project2Id := idFromProject(t, project2)
		if project2Id != project1Id+1 {
			t.Fatalf("expected second id to be %d, but got %d", project1Id+1, project2Id)
		}
	})

	t.Run("itIncrementsProjectIdForASpecificTeamAndWorkTypePair", func(t *testing.T) {
		repo := newSqliteRepository(t, testDsn)

		team1 := model.Team{Name: "FOR"}
		workType1 := model.WorkType{Name: "MA"}
		_, err := repo.CreateNewProject(team1, workType1)
		if err != nil {
			t.Fatal(err)
		}
		project2 := createNewProject(t, repo, team1, workType1)
		if err != nil {
			t.Fatal(err)
		}

		workType2 := model.WorkType{Name: "ES"}
		project3 := createNewProject(t, repo, team1, workType2)
		if err != nil {
			t.Fatal(err)
		}

		p2Id := idFromProject(t, project2)
		p3Id := idFromProject(t, project3)
		if p3Id >= p2Id {
			t.Fatalf("expected id count to reset when changing the worktype")
		}

		team2 := model.Team{Name: "ANA"}
		project4 := createNewProject(t, repo, team2, workType1)
		p4Id := idFromProject(t, project4)

		if p4Id >= p2Id {
			t.Fatalf("expected id count to reset when changing the team")
		}
	})
}

func idFromProject(t *testing.T, project model.Project) int64 {
	projectIdStr := project.Name[len(project.Name)-1:]
	projectId, err := strconv.ParseInt(projectIdStr, 10, 64)
	if err != nil {
		t.Fatalf("unexpected failure. %s", err)
	}
	return projectId
}

func createNewProject(
	t *testing.T,
	repo PngRepository,
	team model.Team,
	workType model.WorkType,
) model.Project {
	project, err := repo.CreateNewProject(team, workType)
	if err != nil {
		t.Fatalf("unexpected failure. %s", err)
	}
	return project
}

func newSqliteRepository(t *testing.T, dsn string) *SqlitePngRepository {
	repo, err := NewSqlitePngRepository(dsn)
	if err != nil {
		t.Fatal(err)
	}
	return repo
}
