package utils

// Result for request
type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// SuccessResult Success result
func SuccessResult(data interface{}) Result {
		return Result{Status: 0, Data: data}
}

// FailResult Fail result
func FailResult(status int, err error) Result {
		return Result{Status: status, Message: err.Error()}
}

// NewResult New a result
func NewResult(data interface{}, err error) Result {
	if err == nil {
		return Result{Status: 0, Data: data}
	} else {
		return Result{Status: 1, Data: data, Message: err.Error()}
	}
}

// NewEmptyResult New a empty reuslt.
func NewEmptyResult(err error) Result {
	if err == nil {
		return Result{Status: 0}
	} else {
		return Result{Status: 1, Message: err.Error()}
	}
}
