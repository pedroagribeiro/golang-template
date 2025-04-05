package server

import (
	"net/http"
	"template/config"
	app "template/core/application"
	"template/core/db"
	"template/core/log"
	"template/docs"

	traditionalLog "log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getDatabaseConnection(cfg *config.Config, application *app.Application) {
	db := db.CreateDb(cfg)
	engine := db.GetEngine()
	application.Db = engine
	log.Info("[getDatabaseConnection]: Initiated connection with the PostgreSQL database")
}

func configureWebRouter(application *app.Application) {
	application.Echo = echo.New()
	application.Router = app.CreateRouter(application)
}

func configureWebRoutes(application *app.Application) {
	docs.BindRoutes(application.Router)
}

func initEntityServices(application *app.Application) {
}

func runAppInitializers(application *app.Application, initializers []app.InitializerHandlerFunc) {
	for _, h := range initializers {
		h(application)
	}
}

func startWebServer(cfg *config.Config) {
	errorChannel := make(chan error, 1)
	requestIdConfig := middleware.DefaultRequestIDConfig
	requestIdConfig.Skipper = func(c echo.Context) bool {
		headers := c.Request().Header
		return headers.Get(app.REQUEST_ID_HEADER) != ""
	}
	requestIdConfig.TargetHeader = app.REQUEST_ID_HEADER
	app.MyApp.Echo.Use(middleware.RequestIDWithConfig(requestIdConfig))
	// app.MyApp.Router.UsePrometheus()
	// start HTTP server
	go func(chan error) {
		if err := app.MyApp.Echo.Start(":" + cfg.Server.HttpPort); err != nil {
			log.Errorf("[startWebServer]: %s", err.Error())
			errorChannel <- err
		}
		app.MyApp.Echo.Logger.Debug(app.MyApp.Echo)
	}(errorChannel)
	// Start HTTPS server - if HTTPS port is specified
	go func(chan error) {
		if err := app.MyApp.Echo.StartTLS(":"+cfg.Server.HttpsPort, cfg.CertificatePath, cfg.KeyPath); err != nil {
			log.Errorf("[startWebServer]: %s", err.Error())
			errorChannel <- err
		}
		app.MyApp.Echo.Logger.Debug(app.MyApp.Echo)
		app.MyApp.Echo.Server.TLSConfig.InsecureSkipVerify = true
	}(errorChannel)

	go func() {
		log.Infof("%+v", http.ListenAndServe("localhost:6060", nil))
	}()
	traditionalLog.Fatal(<-errorChannel)
}

func configureWebServer() {
	cfg := config.GetConfig()
	app.MyApp = new(app.Application)
	getDatabaseConnection(cfg, app.MyApp)
	configureWebRouter(app.MyApp)
	configureWebRoutes(app.MyApp)
	initEntityServices(app.MyApp)
	runAppInitializers(app.MyApp, app.AppInitializers)
	startWebServer(cfg)
}

func Init() {
	configureWebServer()
}
