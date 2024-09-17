package errs

type BadRequestError struct {
	Err string
}

func (e *BadRequestError) Error() string {
	return e.Err
}
