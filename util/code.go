package util

var (
	OK                  = &Errno{Code: 0, Msg: "OK"}
	InternalServerError = &Errno{Code: 10001, Msg: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Msg: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Msg: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Msg: "Database error."}
	ErrToken      = &Errno{Code: 20003, Msg: "Error occurred while signing the JSON web token."}

	//user errors 自由定义
	ErrEncrypt           = &Errno{Code: 20101, Msg: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Msg: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Msg: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Msg: "The password was incorrect."}
)
