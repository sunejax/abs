package store

import (
	"github.com/labstack/echo"
)

// ExternalRoutes ...
func Routes(e *echo.Echo) {
	storeAPI := e.Group("/store/")
	{
		storeAPI.POST(":user/cart/allocate", Allocate)
	}
}
