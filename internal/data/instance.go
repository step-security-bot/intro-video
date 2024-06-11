package data

import (
	"database/sql"
	"os"

	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Instance struct {
	Id             int32
	Videos         map[int32]Video
	Configurations map[int32]Configuration
}

type Video struct {
	Id              int32
	Weight          int32
	ConfigurationId int32
	URL             string
}

type Configuration struct {
	Id     int32
	Bubble config.Bubble
	Cta    config.Cta
}

type Store struct {
	DatabaseUrl string
	DriverName  string
}

func NewStore() (Store, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Store{}, err
	}
	url := os.Getenv("DATABASE_URL")

	return Store{DatabaseUrl: url, DriverName: "libsql"}, nil

}

func (s *Store) LoadInstance(id int32) (Instance, error) {
	db, err := sql.Open(s.DriverName, s.DatabaseUrl)
	if err != nil {
		return Instance{}, err
	}

	rows, err := db.Query(`
		SELECT
			videos.id,
			videos.weight,
			videos.url,
			videos.configuration_id
		FROM instances
		JOIN videos ON videos.instance_id = instances.id
		WHERE instances.id = $1;
		`,
		id,
	)
	defer rows.Close()

	instance := Instance{Id: id, Videos: map[int32]Video{}, Configurations: map[int32]Configuration{}}

	for rows.Next() {
		var video Video

		if err := rows.Scan(
			&video.Id,
			&video.Weight,
			&video.URL,
			&video.ConfigurationId,
		); err != nil {
			return Instance{}, err
		}

		instance.Videos[video.Id] = video
	}

	rows, err = db.Query(`
		SELECT DISTINCT
			config.id,
			config.bubble_enabled,
			config.bubble_text_content,
			config.bubble_type,
			config.cta_enabled,
			config.cta_text_content,
			config.cta_type
		FROM instances
		JOIN videos ON videos.instance_id = instances.id
		JOIN configurations as config ON videos.configuration_id = config.id
		WHERE instances.id = $1;
		`,
		id,
	)
	defer rows.Close()

	for rows.Next() {
		var configuration Configuration

		configuration.Bubble = config.Bubble{}
		configuration.Cta = config.Cta{}

		if err := rows.Scan(
			&configuration.Id,
			&configuration.Bubble.Enabled,
			&configuration.Bubble.TextContent,
			&configuration.Bubble.Type,
			&configuration.Cta.Enabled,
			&configuration.Cta.TextContent,
			&configuration.Cta.Type,
		); err != nil {
			return Instance{}, err
		}
		instance.Configurations[configuration.Id] = configuration
	}

	return instance, nil
}

func (s *Store) SaveInstance(id int32, instance map[int32]Video) error {
	return nil
}
