package utils

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

// ResultStatus error code.
type ResultStatus int

const (
	// Success success.
	Success ResultStatus = iota
	// Unknown unknown error.
	Unknown
	// NeedLogin need login.
	NeedLogin
	// UserNotExist user not exist.
	UserNotExist
	// PasswordError password error.
	PasswordError
	// NeedAdmin need admin user.
	NeedAdmin
	// LikeMoreTimes like more times.
	LikeMoreTimes
	// DislikeMoreTimes dislike more times.
	DislikeMoreTimes
)

var (
	// ErrNeedLogin need login error.
	ErrNeedLogin = errors.New("need login")
	// ErrUserNotExist user not exist errror.
	ErrUserNotExist = errors.New("user not exist")
	// ErrPasswordError password error
	ErrPasswordError = errors.New("password errror")
	// ErrNeedAdmin need admin user.
	ErrNeedAdmin = errors.New("need admin user")
	// ErrLikeMoreTimes like more times.
	ErrLikeMoreTimes = errors.New("like more times")
	// ErrDislikeMoreTimes dislike more times.
	ErrDislikeMoreTimes = errors.New("dislike more times")
)

var errorMap = map[error]ResultStatus{
	ErrNeedLogin: NeedLogin,
	ErrUserNotExist: UserNotExist,
	ErrPasswordError: PasswordError,
	ErrNeedAdmin: NeedAdmin,
	ErrLikeMoreTimes: LikeMoreTimes,
	ErrDislikeMoreTimes: DislikeMoreTimes,
}

// Result for request
type Result struct {
	Status  ResultStatus `json:"status"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message"`
}

// SuccessResult Success result
func SuccessResult(data interface{}) Result {
	return Result{Status: Success, Data: data}
}

// FailResult Fail result
func FailResult(err error) Result {
	status := Unknown
	if _, ok := errorMap[err]; ok {
		status = errorMap[err]
	}
	return Result{Status: status, Message: err.Error()}
}

// NewResult New a result
func NewResult(data interface{}, err error) Result {
	if err == nil || err == orm.ErrNoRows {
		return Result{Status: Success, Data: data}
	}

	status := Unknown
	if _, ok := errorMap[err]; ok {
		status = errorMap[err]
	}
	return Result{Status: status, Data: data, Message: err.Error()}
}

// NewEmptyResult New a empty reuslt.
func NewEmptyResult(err error) Result {
	if err == nil {
		return Result{Status: Success}
	}

	status := Unknown
	if _, ok := errorMap[err]; ok {
		status = errorMap[err]
	}
	return Result{Status: status, Message: err.Error()}
}
