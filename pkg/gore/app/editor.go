package app

type Editor interface {
	EditFilenames(filenames []string) ([]string, error)
}
