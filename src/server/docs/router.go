package docs

import (
	app "template/core/application"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func BindRoutes(router *app.Router) {
	app.MyApp.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
