package errmessages

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
	ErrNotFound        = errors.New("not found")
)

type Error struct {
	appErr   error
	svcError error
}

func (e *Error) AppError() error {
	return e.appErr
}

func (e *Error) SvcError() error {
	return e.svcError
}

func NewError(svcError, appError error) Error {
	return Error{
		appErr:   appError,
		svcError: svcError,
	}
}

func (e *Error) Error() string {
	return errors.Join(e.appErr, e.svcError).Error()
}
