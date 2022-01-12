package apiclient

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int
}

func NewError(status int, code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Status:  status,
	}
}
