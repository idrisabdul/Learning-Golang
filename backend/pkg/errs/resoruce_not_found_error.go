package errs

type ResourceNotFoundError struct {
	Err string
}

func (e *ResourceNotFoundError) Error() string {
	return e.Err
}
