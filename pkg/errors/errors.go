package errors

type IError struct {
	ErrorType string
	Message   error
}

func (e *IError) Error() string {
	return e.ErrorType + ": " + e.Message.Error()
}

func NewError(errorType string, message error) *IError {
	return &IError{ErrorType: errorType, Message: message}
}
