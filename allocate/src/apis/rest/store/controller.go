package store

import (
	"github.com/labstack/echo"
	"net/http"
	shopAllocateModule "repo.abs.com/hiring/allocate/src/modules/allocate"
)

// Allocate ...
func Allocate(c echo.Context) error {
	// Forms Request (Raw to map)
	req, err := bindJSON(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
	}

	req["user"] = c.Param("user")

	// Calls BL
	res, err := shopAllocateModule.Allocate(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
	}

	// Returns
	return c.JSON(http.StatusOK, res)
}

func bindJSON(c echo.Context) (result map[string]interface{}, err error) {
	// Binds JSON data to a map
	requestBody := make(map[string]interface{})
	err = c.Bind(&requestBody)
	if err != nil {
		return nil, err
	}

	return requestBody, nil
}
