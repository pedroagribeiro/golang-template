package main

import (
	"template/docs"
	"template/server"
)

// @title			Application name
// @version			0.0
// @description		Description of the application

// @contact.name	Developers and maintainers name
// @contact.email	Developers and maintainers email contacts

// @host			localhost
// @port			9000
// @BasePath		/
func main() {
	docs.SwaggerInfo.Host = "localhost:9000"
	server.Init()
}
