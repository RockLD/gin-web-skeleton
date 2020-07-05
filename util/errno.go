package util

import "fmt"

type Errno struct {
	Code int
	Msg  string
}

func (err Errno) Error() string {
	return err.Msg
}

type Err struct {
	Code int
	Msg  string
	Err  error
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Msg: errno.Msg, Err: err}
}

func (err *Err) Addf(format string, args ...interface{}) *Err {
	err.Msg += " " + fmt.Sprintf(format, args)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, msg: %s, error: %s", err.Code, err.Msg, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Errno:
		return typed.Code, typed.Msg
	default:
	}
	return InternalServerError.Code, err.Error()
}
