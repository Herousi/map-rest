package res

import (
	"fmt"
	"go.uber.org/zap"
)

type BgRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (b BgRes) ErrorDetail(err error) BgRes {
	if err != nil {
		b.Data = err.Error()
	}
	return b
}

func (b BgRes) ErrorDes(errDes string) BgRes {
	b.Data = errDes
	return b
}

func (b BgRes) Error() string {
	return fmt.Sprintf("code:%d Message:%s", b.Code, b.Msg)
}

func DecodeErr(err error, data interface{}) (int, string, interface{}) {
	if err == nil {
		return OK.Code, OK.Msg, data
	}

	switch typed := err.(type) {
	case BgRes:
		return typed.Code, typed.Msg, typed.Data
	case *BgRes:
		return typed.Code, typed.Msg, typed.Data
	}
	zap.L().Error("响应错误信息被拦截", zap.Error(err))
	return InternalServerError.Code, InternalServerError.Msg, err.Error()
}
