package exception

import "net/http"

// interface for error
type MessageErr interface {
	Status() int
	Error() string
}

type ErrData struct {
	ErrMessage interface{} `json:"message"`
	ErrStatus  int         `json:"status"`
	ErrError   string      `json:"error"`
}

func (e *ErrData) Status() int {
	return e.ErrStatus
}

func (e *ErrData) Error() string {
	return e.ErrError
}

func BadRequest(message string) MessageErr {
	return &ErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}

func InternalServerError(message string) MessageErr {
	return &ErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func UnprocessableEntity(message []string) MessageErr {
	return &ErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "UNPROCESSABLE_ENTITY",
	}
}

func NotFound(message string) MessageErr {
	return &ErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func Unauthorized(message string) MessageErr {
	return &ErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "UNAUTHORIZED",
	}
}
