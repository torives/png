package repo

import (
	"fmt"
	"os"
	"testing"
)

var testDbPath = "./test/data/test.sqlite"

func SetupError(err error) error {
	return fmt.Errorf("failed to set up test: %w", err)
}

//FIXME: tests could be performed with the memory database

// Open
func TestItCreatesADatabaseIfNoneExists(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "repo_test")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.Remove(tmpDir)

	db, err := Open(tmpDir + "/test.sqlite")
	if err != nil {
		t.Fatalf(fmt.Sprintf("failed to create a database. %s", err))
	}

	if db == nil {
		t.Fatal("database should not be nil")
	}
}

func TestItOpensAnExistingDatabase(t *testing.T) {
	db, err := Open(testDbPath)
	if err != nil {
		t.Fatal("failed to open existing database")
	}
	if db == nil {
		t.Fatal("database should not be nil")
	}
}

// Close
func TestItClosesTheDatabase(t *testing.T) {}
func TestItFailsToQueryTheDatabaseAfterClosing(t *testing.T) {
	db, err := Open(testDbPath)
	if err != nil {
		t.Fatal(SetupError(err))
	}
	db.Close()

	//TODO: find out how to compare errors during testing
	err = db.AddCategory("treta")
	if err == nil {
		t.Fatalf("error should not be nil. %s", err)
	}
}

// New Category
func TestItAddsANewCategory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	err = db.AddCategory("test")
	if err != nil {
		t.Fatalf("should have created category. %s", err)
	}
}

func TestItFailsToAddNewCategoryWithSameName(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	category := "test"
	err = db.AddCategory(category)
	if err != nil {
		t.Fatalf("should have created category. %s", err)
	}
	err = db.AddCategory(category)
	if err == nil {
		t.Fatalf("should have failed to create a duplicate category")
	}
}

// Next Project Id
func TestItGetsProjectId(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	category := "test"
	err = db.AddCategory(category)
	if err != nil {
		t.Fatalf("should have created category. %s", err)
	}

	_, err = db.NextProjectId(category)
	if err != nil {
		t.Fatalf("should have returned next project's id. %s", err)
	}
}

func TestItAlwaysReturnsAGreaterId(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	category := "test"
	err = db.AddCategory(category)
	if err != nil {
		t.Fatalf("should have created category. %s", err)
	}

	firstId, err := db.NextProjectId(category)
	if err != nil {
		t.Fatalf("should have returned next project's id. %s", err)
	}

	secondId, err := db.NextProjectId(category)
	if err != nil {
		t.Fatalf("should have returned next project's id. %s", err)
	}

	if firstId >= secondId {
		t.Fatalf(
			"NextProjectId should always return bigger ids. Got %d after %d",
			secondId, firstId,
		)
	}
}
func TestProjectIdsAreAlwaysIncrementedByOne(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	category := "test"
	err = db.AddCategory(category)
	if err != nil {
		t.Fatalf("should have created category. %s", err)
	}

	firstId, err := db.NextProjectId(category)
	if err != nil {
		t.Fatalf("should have returned next project's id. %s", err)
	}

	secondId, err := db.NextProjectId(category)
	if err != nil {
		t.Fatalf("should have returned next project's id. %s", err)
	}

	if firstId+1 != secondId {
		t.Fatalf(
			"NextProjectId should return last id increased by one. Got %d after %d",
			secondId, firstId,
		)
	}
}

func TestItFailsToGetProjectIdIfCategoryDoesNotExist(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_repo")
	if err != nil {
		t.Fatal(SetupError(err))
	}
	defer os.RemoveAll(tmpDir)

	db, err := Open(tmpDir + "/new.sqlite")
	if err != nil {
		t.Fatal(SetupError(err))
	}

	category := "test"
	_, err = db.NextProjectId(category)
	if err == nil {
		t.Fatal("should have failed to get id from non-existing category")
	}
}
