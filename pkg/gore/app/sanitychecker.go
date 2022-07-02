package app

type SanityChecker interface {
	Verify(filesBefore []string, filesAfter []string) error
}

func NewSanityChecker() SanityChecker {
	return &sanityChecker{}
}

type sanityChecker struct{}

func (c *sanityChecker) Verify(filesBefore []string, filesAfter []string) error {
	// TODO implement me
	return nil
}
