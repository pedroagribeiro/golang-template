package application

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"template/core/db"
	errors "template/core/error"

	"github.com/labstack/echo/v4"
)

type IContext interface {
	Db() db.IDbConnection
	Request() IContextRequest
	Response() IContextRequest
	Set(key string, data interface{})
	Get(key string) interface{}
}

func (c RouterContext) Db() db.IDbConnection {
	return c.app.Db
}

func (c RouterContext) Request() IContextRequest {
	return RouterContextRequest{ctx: &c}
}

func (c RouterContext) Response() IContextResponse {
	return RouterContextResponse{ctx: &c}
}

func (c RouterContext) Set(key string, data interface{}) {
	c.echo.Set(key, data)
}

func (c RouterContext) Get(key string) interface{} {
	return c.echo.Get(key)
}

type IContextRequest interface {
	Body(data any) error
	Query(data any) error
	Host() string
	Param() *RouterContextRequestParam
	GetParam(key string) string
	Headers() http.Header
	GetMultipartForm() (*multipart.Form, error)
	GetFormFile(string) (*multipart.FileHeader, error)
}

func (r RouterContextRequest) Body(data interface{}) error {
	return r.ctx.echo.Bind(data)
}

func (r RouterContextRequest) Query(data interface{}) error {
	return (&echo.DefaultBinder{}).BindQueryParams(r.ctx.echo, data)
}

func (r RouterContextRequest) Host() string {
	return r.ctx.echo.Request().Host
}

func (r RouterContextRequest) Param() *RouterContextRequestParam {
	return &RouterContextRequestParam{ctx: r.ctx}
}

func (r RouterContextRequest) GetParam(key string) string {
	return r.ctx.echo.Param(key)
}

func (r RouterContextRequest) Headers() http.Header {
	return r.ctx.echo.Request().Header
}

func (r RouterContextRequest) GetMultipartForm() (*multipart.Form, error) {
	return r.ctx.echo.MultipartForm()
}

func (r RouterContextRequest) GetFormFile(fileName string) (*multipart.FileHeader, error) {
	return r.ctx.echo.FormFile(fileName)
}

type IContextResponse interface {
	ErrorFromGenericError(errors.GenericError) error
	Error(httpStatusCode int, code int, cause string, message string, status string, beans ...interface{}) error
	String(httpStatusCode int, body string) error
	JSON(httpStatusCode int, i any) error
	File(httpStatusCode int, filePath string) error
	Headers() http.Header
}

func (r RouterContextResponse) ErrorFromGenericError(generic_error errors.GenericError) error {
	return r.ctx.echo.JSON(generic_error.RestError.ErrorCode, generic_error.RestError)
}

func (r RouterContextResponse) Error(httpStatusCode int, code int, cause string, message string, status string, beans ...interface{}) error {
	return r.ctx.echo.JSON(httpStatusCode, errors.RouterResponseError{ErrorCode: code, Cause: cause, Message: fmt.Sprintf(message, beans...), Status: status})
}

func (r RouterContextResponse) String(httpStatusCode int, body string) error {
	return r.ctx.echo.String(httpStatusCode, body)
}

func (r RouterContextResponse) JSON(httpStatusCode int, i any) error {
	return r.ctx.echo.JSON(httpStatusCode, i)
}

func (r RouterContextResponse) File(httpStatusCode int, filePath string) error {
	return r.ctx.echo.File(filePath)
}

func (r RouterContextResponse) Headers() http.Header {
	return r.ctx.echo.Response().Header()
}
