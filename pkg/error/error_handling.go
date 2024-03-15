package errHandling

import (
	"errors"
	errmessages "golang_testing_grpc/pkg/err_messages"
	"net/http"
)

type ApiError struct {
	Status  int
	Message string
}

func FromError(err error) ApiError {
	var apiError ApiError
	var svcError *errmessages.Error
	if errors.As(err, &svcError) {
		apiError.Message = svcError.AppError().Error()
		svcError := svcError.SvcError()
		switch svcError {
		case errmessages.ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case errmessages.ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		case errmessages.ErrNotFound:
			apiError.Status = http.StatusNotFound
		}
	}
	return apiError
}
