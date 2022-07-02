package app

type SanityChecker interface {
	Verify(filesBefore, filesAfter []string) error
}

func NewSanityChecker() SanityChecker {
	return &sanityChecker{}
}

type sanityChecker struct{}

func (c *sanityChecker) Verify(filesBefore, filesAfter []string) error {
	return nil
}
