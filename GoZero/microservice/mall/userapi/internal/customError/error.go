package customError

var ParamsError = NewBizError(500, "参数错误")

// 定义一个业务错误
type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewBizError(code int, msg string) *BizError {
	return &BizError{
		code, msg,
	}
}

func (e BizError) Error() string {
	return e.Msg
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e BizError) Data() *ErrorResponse {
	return &ErrorResponse{
		e.Code, e.Msg,
	}
}
