package errs

type ForbiddenError struct {
	Err string
}

func (e *ForbiddenError) Error() string {
	return e.Err
}
