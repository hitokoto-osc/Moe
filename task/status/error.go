package status

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// GenStatusRequestFailureError 定义了统计请求时产生异常的错误类型
type GenStatusRequestFailureError struct {
	Code         int             // HTTP Response Code, or API Error code
	Detail       string          // Error Message
	ResponseData *resty.Response // Response Data
	Stack        []byte          // Call Stack
}

func (e *GenStatusRequestFailureError) Error() string {
	return fmt.Sprintf("[request] request failed. %v:%v", e.Code, e.Detail)
}
