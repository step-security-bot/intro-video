package data

import (
	"database/sql"

	"github.com/crocoder-dev/intro-video/internal"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Video struct {
	id int32
	weight int32
	internal.ProcessableFileProps
}

func LoadInstance(id int32) (map[int32]Video, error) {
	url := ""
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(
		`SELECT
			video.id,
			instance.weight,
			conf.id,
			conf.bubble_enabled,
			conf.bubble_text,
			conf.cta_enabled,
			conf.cta_text
		FROM instance
		JOIN video ON instance.video_id = video.id


		`
	)
	return nil, nil
}

func SaveInstance(id int32, instance map[int32]Video) error {
	return nil
}
