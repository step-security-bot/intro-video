package internal

import "github.com/crocoder-dev/intro-video/internal/config"

type ProcessableFileProps struct {
	config.Bubble
	config.Cta
	config.Video
}

type ProcessableFile interface {
	Process(props ProcessableFileProps) (string, error)
}
