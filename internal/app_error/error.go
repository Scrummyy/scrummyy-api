package apperror

type (
	AppErrorInterface interface {
		error
		GetHttpCode() int
		GetError() error
		GetErrorCode() string
	}

	AppError struct {
		super     error
		httpCode  int
		errorCode string
	}
)

// New returns an instance of AppError.
func New(err error, code int) AppErrorInterface {
	return &AppError{
		super:    err,
		httpCode: code,
	}
}

// New returns an instance of AppError.
func NewWithErrorCode(err error, code int, errCode string) AppErrorInterface {
	return &AppError{
		super:     err,
		httpCode:  code,
		errorCode: errCode,
	}
}

// Error returns the error string.
func (ae *AppError) Error() string {
	return ae.super.Error()
}

// GetHttpCode returns the http code if any.
func (ae *AppError) GetHttpCode() int {
	return ae.httpCode
}

// GetError return the actual error.
func (ae *AppError) GetError() error {
	return ae.super
}

// GetError return the actual error.
func (ae *AppError) GetErrorCode() string {
	return ae.errorCode
}
