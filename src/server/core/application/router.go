package application

import (
	"net/http"
	"template/core/oswrap"

	"github.com/labstack/echo-contrib/echoprometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const PrometheusMetricsEndpoints = "/prometheus/metrics"

type Router struct {
	echo           *echo.Echo
	app            *Application
	protectedGroup *echo.Group
}

type RouterContext struct {
	app  *Application
	echo echo.Context
	http *http.Request
}

type RouterContextRequest struct {
	ctx *RouterContext
}

type RouterContextRequestParam struct {
	ctx *RouterContext
}

type RouterContextResponse struct {
	ctx *RouterContext
}

type List[T interface{}] struct {
	Hits  []T `json:"hits"`
	Total int `json:"total"`
}

type HandlerRouterFunc func(RouterContext) error
type MiddlewareFunc echo.MiddlewareFunc

const REQUEST_ID_HEADER = "UNIQUE_ID"

func GetRequestId(rc RouterContext) string {
	headers := rc.Request().Headers()
	return headers.Get(REQUEST_ID_HEADER)
}

func CreateRouter(app *Application) *Router {
	output := Router{}
	output.app = app
	output.echo = app.Echo
	output.protectedGroup = output.echo.Group("")
	output.protectedGroup.Use(echojwt.WithConfig(Config))
	return &output
}

func CreateRouterContext(c echo.Context, r *Router) RouterContext {
	context := RouterContext{echo: c}
	context.app = r.app
	context.http = c.Request()
	return context
}

func CreateRequestRoutineContext(c echo.Context, r *Router) (routerContext RouterContext) {
	routerContext = CreateRouterContext(c, r)
	requestId := GetRequestId(routerContext)
	routineContext := oswrap.GetRoutineContext()
	oswrap.RoutineContextWithValue(routineContext, "requestId", requestId)
	return
}

func (r *Router) GET(path string, handler HandlerRouterFunc) *Router {
	r.protectedGroup.GET(path, ProtectedWithJwtHandler(r, handler))
	return r
}

func (r *Router) DELETE(path string, handler HandlerRouterFunc) *Router {
	r.protectedGroup.DELETE(path, ProtectedWithJwtHandler(r, handler))
	return r
}

func (r *Router) POST(path string, handler HandlerRouterFunc) *Router {
	r.protectedGroup.POST(path, ProtectedWithJwtHandler(r, handler))
	return r
}

func (r *Router) PUT(path string, handler HandlerRouterFunc) *Router {
	r.protectedGroup.PUT(path, ProtectedWithJwtHandler(r, handler))
	return r
}

func (r *Router) UnprotectedPOST(path string, handler HandlerRouterFunc) *Router {
	r.echo.POST(path, UnprotectedHandler(r, handler))
	return r
}

func UnprotectedHandler(r *Router, handler HandlerRouterFunc) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		return HandleRequest(r, c, handler)
	}
}

func HandleRequest(r *Router, c echo.Context, handler HandlerRouterFunc) (err error) {
	routerContext := CreateRequestRoutineContext(c, r)
	defer oswrap.DeleteRoutineContext()
	err = handler(routerContext)
	return
}

func (r *Router) UsePrometheus() {
	r.echo.Use(echoprometheus.NewMiddleware("incidenthub_go"))
	r.echo.GET(PrometheusMetricsEndpoints, echoprometheus.NewHandler())
}
