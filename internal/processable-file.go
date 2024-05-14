package internal

import "github.com/crocoder-dev/intro-video/internal/config"

type ProcessableFileProps struct {
	config.Bubble
	config.Cta
	URL string
}

type ProcessableFile interface {
	Process(props ProcessableFileProps) (string, error)
}
