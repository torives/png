package repository

import (
	"fmt"
	"testing"

	"github.com/torives/png/model"
)

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
			t.Fatalf("expected failure. %s", err)
		}
	})
}
