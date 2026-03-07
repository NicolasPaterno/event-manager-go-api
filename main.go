package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicao/minimal-goapi/db"
	"github.com/nicao/minimal-goapi/routes"
	_ "github.com/nicao/minimal-goapi/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Minimal GoAPI
// @version         1.0
// @description     API de eventos com autenticação JWT
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}
