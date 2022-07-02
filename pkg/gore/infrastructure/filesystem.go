package infrastructure

import (
	"os"
	"path/filepath"

	"github.com/m3tro1d/gore/pkg/gore/app"
)

func NewFilesystem() app.Filesystem {
	return &filesystem{}
}

type filesystem struct{}

func (s *filesystem) RenameFiles(directory string, filenames, newFilenames []string) error {
	for i := 0; i < len(filenames); i++ {
		path := filepath.Join(directory, filenames[i])
		newPath := filepath.Join(directory, newFilenames[i])

		err := os.Rename(path, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}
