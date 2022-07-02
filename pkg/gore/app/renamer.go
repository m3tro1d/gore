package app

import (
	"os"
	"path/filepath"
)

func NewRenamer(fileLister FileLister, editor Editor, sanityChecker SanityChecker) *Renamer {
	return &Renamer{
		fileLister:    fileLister,
		editor:        editor,
		sanityChecker: sanityChecker,
	}
}

type Renamer struct {
	fileLister    FileLister
	editor        Editor
	sanityChecker SanityChecker
}

func (r *Renamer) Rename(directory string) error {
	filenames, err := r.fileLister.ListFiles(directory)
	if err != nil {
		return err
	}

	newFilenames, err := r.editor.EditFilenames(filenames)
	if err != nil {
		return err
	}

	err = r.sanityChecker.Verify(filenames, newFilenames)
	if err != nil {
		return err
	}

	return r.renameFiles(directory, filenames, newFilenames)
}

func (r *Renamer) renameFiles(directory string, filenames, newFilenames []string) error {
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
