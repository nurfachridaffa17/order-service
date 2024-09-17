package routes

import (
	"net/http"
	"order-service/config"
	"order-service/internal/api/http/routes/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(g *echo.Group, db *gorm.DB) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+config.NewEnv().GetString("APP")+"! version "+config.NewEnv().GetString("VERSION")+" in mode "+config.NewEnv().GetString("ENV"))
	})

	handler.NewOrderHandler(db).Route(g.Group("/order"))
}
