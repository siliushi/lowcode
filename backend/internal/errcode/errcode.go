package errcode

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ErrCode int

const (
	Success               ErrCode = 0
	ErrNone               ErrCode = Success
	errBase               ErrCode = 10000
	ErrInvalidParam       ErrCode = errBase + 1
	ErrInvalidToken       ErrCode = errBase + 2
	errAuthBase           ErrCode = 10100
	ErrAuthInvalidParam   ErrCode = errAuthBase + 1
	ErrAuthPasswdNotMatch ErrCode = errAuthBase + 2
	ErrAuthNameNotExit    ErrCode = errAuthBase + 3
	ErrAuthDB             ErrCode = errAuthBase + 4
	errBooksBase          ErrCode = 10200
	ErrBooksNotExit       ErrCode = errBooksBase + 1
	ErrBooksDeleteFail    ErrCode = errBooksBase + 2
	ErrBooksUpdateFail    ErrCode = errBooksBase + 3
)

var errCodeMap = map[ErrCode]string{
	ErrNone:         "Success",
	ErrInvalidParam: "参数错误",
	ErrInvalidToken: "认证失败",
	// Auth
	ErrAuthInvalidParam:   "用户名或者密码为空",
	ErrAuthPasswdNotMatch: "密码不正确",
	ErrAuthNameNotExit:    "用户名不存在",
	ErrAuthDB:             "数据库错误",
	// Books
	ErrBooksNotExit:    "书籍不存在",
	ErrBooksDeleteFail: "书籍删除失败",
	ErrBooksUpdateFail: "书籍更新失败",
}

type CodeError struct {
	stackErr error
	Code     ErrCode
	Msg      string
}

type CodeErrorResponse struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
}

func NewCodeError(code ErrCode, err error) error {
	return &CodeError{Code: code, Msg: errCodeMap[code], stackErr: errors.WithStack(err)}
}

func NewDefaultError(err error) error {
	return NewCodeError(ErrInvalidParam, err)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// HandleError 自定义错误
func HandleError(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
		logx.Errorf("%+v", e.stackErr)
		return http.StatusOK, e.Data()
	default:
		logx.Errorf("%+v", err)
		return http.StatusOK, &CodeErrorResponse{
			Code: ErrInvalidParam,
			Msg:  errCodeMap[ErrInvalidParam],
		}
	}
}

// OKResponse : Success Response for HTTP API
type OKResponse struct {
	Code ErrCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(data interface{}) *OKResponse {
	return &OKResponse{
		Code: ErrNone,
		Msg:  errCodeMap[ErrNone],
		Data: data,
	}
}

func ValidateErr(err error) *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: ErrInvalidParam,
		Msg:  err.Error(),
	}
}
