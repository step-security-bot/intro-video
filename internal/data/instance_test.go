package data_test

import (
	"database/sql"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/crocoder-dev/intro-video/internal/data"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func getMigrationSchemas() ([]string, error) {
	migrationsPath := filepath.Join("..", "..", "db", "migrations")

	var schemaFiles []string

	err := filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".up.sql") {
			schemaFiles = append(schemaFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(schemaFiles)

	return schemaFiles, nil
}

func TestLoadInstance(t *testing.T) {

	file, err := os.CreateTemp("", "test*.db")

	if err != nil {
		t.Fatalf("failed to create database file: %v", err)
	}

	defer os.Remove(file.Name())

	db, err := sql.Open("sqlite3", file.Name())

	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	schemaFiles, err := getMigrationSchemas()

	if err != nil {
		t.Fatalf("failed to read schema files: %v", err)
	}

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

	_, err = db.Exec(`
		INSERT INTO instances (id) VALUES (1);

		INSERT INTO configurations (id, bubble_enabled, bubble_text_content, bubble_type, cta_enabled, cta_text_content, cta_type)
		VALUES (1, 1, "bubble text", "default", 1, "cta text", "default");

		INSERT INTO videos (id, instance_id, configuration_id, weight, url)
		VaLUES (1, 1, 1, 100, "url");
	`)

	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	store := data.Store{DatabaseUrl: file.Name(), DriverName: "sqlite3"}

	instance, err := store.LoadInstance(1)

	expected := data.Instance{
		Id: 1,
		Videos: map[int32]data.Video{
			1: {
				Id:              1,
				Weight:          100,
				ConfigurationId: 1,
				URL:             "url",
			},
		},
		Configurations: map[int32]data.Configuration{
			1: {
				Id: 1,
				Bubble: config.Bubble{
					Enabled:     true,
					TextContent: "bubble text",
					Type:        "default",
				},
				Cta: config.Cta{
					Enabled:     true,
					TextContent: "cta text",
					Type:        "default",
				},
			},
		},
	}

	if len(instance.Videos) != len(expected.Videos) {
		t.Fatalf("Expected %d videos, got %d", len(expected.Videos), len(instance.Videos))
	}

	if len(instance.Configurations) != len(expected.Configurations) {
		t.Fatalf("Expected %d configurations, got %d", len(expected.Configurations), len(instance.Configurations))
	}

	for id, video := range expected.Videos {
		if v, ok := instance.Videos[id]; !ok || v != video {
			t.Fatalf("Video with id %d not found or does not match expected", id)
		}
	}

	for id, config := range expected.Configurations {
		if c, ok := instance.Configurations[id]; !ok || c != config {
			t.Fatalf("Configuration with id %d not found or does not match expected", id)
		}
	}

}
