package store

import (
	"github.com/labstack/echo"
	"net/http"
	shopGetQuantityModule "repo.abs.com/hiring/getQuantity/src/modules/getquantity"
)

// GetQuantity ...
func GetQuantity(c echo.Context) error {
	// Forms Request (Raw to map)
	req := make(map[string]interface{})

	req["productCode"] = c.Param("productCode")

	// Calls BL
	res, err := shopGetQuantityModule.GetQuantity(c.Request().Context(), req)

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
