package pkg

import (
	"fmt"
)

type CollectError struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status int    `json:"-"`
	Raw    error  `json:"-"`
}

var ErrorPkg = new(CollectError)

func (e *CollectError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func (e *CollectError) Unwrap() error {
	return e.Raw
}
func New(code int, msg string, statusCode int) *CollectError {
	err := new(CollectError)
	err.Code = code
	err.Msg = msg
	err.Status = statusCode
	return err
}
func (e *CollectError) WithCause(err error) *CollectError {
	newErr := *e
	newErr.Raw = err
	return &newErr
}
