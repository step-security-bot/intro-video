package data_test

import (
	"database/sql"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func TestLoadInstance(t *testing.T) {

	file, err := os.CreateTemp("", "test*.db")

	if err != nil {
		t.Fatalf("failed to create database file: %v", err)
	}

	defer os.Remove(file.Name())

	migrationsPath := filepath.Join("..", "..", "db", "migrations")

	var schemaFiles []string

	err = filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".up.sql") {
			schemaFiles = append(schemaFiles, path)
		}

		return nil
	})

	if err != nil {
		t.Fatalf("failed to read schema files: %v", err)
	}

	sort.Strings(schemaFiles)

	db, err := sql.Open("sqlite3", file.Name())

	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	for _, schemaFile := range schemaFiles {
		schema, err := os.ReadFile(schemaFile)

		if err != nil {
			t.Fatalf("failed to read schema file %s: %v", schemaFile, err)
		}

		_, err = db.Exec(string(schema))

		if err != nil {
			t.Fatalf("failed to execute schema %s: %v", schemaFile, err)
		}
	}


	t.Fatalf("Not implemented")

}
