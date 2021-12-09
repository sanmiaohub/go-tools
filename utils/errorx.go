package utils

type XError struct {
	code    int
	message string
}

func NewError(c int, s string) *XError {
	return &XError{code: c, message: s}
}

func (e XError) Code() int {
	return e.code
}
func (e XError) Error() string {
	return e.message
}
