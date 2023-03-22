package utils

import (
	"image"
	_ "image/png"
	"io/fs"
)

type Loader struct {
	filesystem fs.FS
}

func NewLoader(filesystem fs.FS) *Loader {
	return &Loader{filesystem}
}

func (l *Loader) Open(path string) (fs.File, error) {
	return l.filesystem.Open(path)
}

func (l *Loader) Image(path string) (image.Image, error) {
	file, err := l.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
