package utils

type ErrorX struct {
	code    int
	message string
}

func NewError(c int, s string) *ErrorX {
	return &ErrorX{code: c, message: s}
}

func (e ErrorX) Code() int {
	return e.code
}
func (e ErrorX) Error() string {
	return e.message
}
