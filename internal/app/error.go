package app

type Code int

type Result struct {
	Status Code
	Error  error
}

const (
	CODE_OK        Code = 0
	CODE_NQ_SIZE   Code = 1
	CODE_NQ_HASH   Code = 2
	CODE_NQ_COUNT  Code = 3
	CODE_NOT_EXIST Code = 4
	CODE_ERROR     Code = 10
)

var (
	OK = Result{CODE_OK, nil}
)

func NewError(code Code, err error) Result {
	return Result{code, err}
}
