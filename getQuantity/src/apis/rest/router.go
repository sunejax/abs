package rest

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"repo.abs.com/hiring/getQuantity/src/apis/rest/store"
	"repo.abs.com/hiring/getQuantity/src/config"
)

// BuildRouter ...
func BuildRouter() {
	// Initialises internal server
	router := echo.New()

	router.GET("/ping", checkPing)
	store.Routes(router)

	err := router.Start(config.GConfig.Server.HTTP.Port)
	if err != nil {
		fmt.Println(err)
	}
}

func checkPing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
