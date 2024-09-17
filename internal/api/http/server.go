package http

import (
	"net/http"
	"order-service/config"
	"order-service/internal/api/http/routes"
	"order-service/internal/pkg/logging"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Auth-Token"},
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	routes.Init(e.Group("/api/v1"), db)
	logging.Log.Info("Server started on port " + config.NewEnv().GetString("PORT"))
	e.Logger.Fatal(e.Start(":" + config.NewEnv().GetString("PORT")))

}
