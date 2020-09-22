package restserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"echo-swaggo-example/pkg/controller"
	"log"
)

// @title Echo-Swagger Example API
// @version 1.0
// @description This is a sample API documentation of Echo-Swagger.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host echo.swagger.io
// @BasePath /v1

type Rest struct {
	key    string
	server *echo.Echo
	Port   string
}

func (r *Rest) StartServer() {
	r.server = echo.New()
	r.server.Use(middleware.Recover())
	r.loadAPI()

	err := r.server.Start(":" + r.Port)
	if err != nil {
		log.Fatal("error starting Echo server")
	}
}

func (r *Rest) loadAPI() {

	c := controller.NewController()

	v1 := r.server.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}

	}
	r.server.GET("/swagger/*", echoSwagger.WrapHandler)
}
