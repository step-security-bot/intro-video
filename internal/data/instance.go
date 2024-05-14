package data

import "github.com/crocoder-dev/intro-video/internal"

type Video struct {
	id int32
	weight int32
	internal.ProcessableFileProps
}

func LoadInstance(id int32) (map[int32]Video, error) {
	return nil, nil
}

func SaveInstance(id int32, instance map[int32]Video) error {
	return nil
}
