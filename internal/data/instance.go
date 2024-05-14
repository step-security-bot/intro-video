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
	rows, err := db.Query(`
		SELECT
			videos.id,
			videos.weight,
			confs.id,
			confs.bubble_enabled,
			confs.bubble_text_content,
			confs.cta_enabled,
			confs.cta_text_content
		FROM instances
		JOIN videos ON videos.instance_id = instances.id
		JOIN configurations as confs ON confs.id = videos.configuration_id;
		WHERE instances.id = $1;
		`,
		id,
	)
	return nil, nil
}

func SaveInstance(id int32, instance map[int32]Video) error {
	return nil
}
