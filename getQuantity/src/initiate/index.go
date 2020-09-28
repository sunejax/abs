package initiate

import (
	"errors"
	"os"
	router "repo.abs.com/hiring/getQuantity/src/apis/rest"
	"repo.abs.com/hiring/getQuantity/src/config"
	"repo.abs.com/hiring/getQuantity/src/dal"
	"repo.abs.com/hiring/getQuantity/src/utils"
	"strings"
)

func Initialise() (err error) {

	// Initialises the config
	err = initialiseConfig()
	if err != nil {
		return errors.New("config init error ->" + err.Error())
	}

	// Initializes DAL
	err = dal.Initialize()
	if err != nil {
		return errors.New("dal init error ->" + err.Error())
	}

	startServers()
	return nil
}

// Initialise ...
func initialiseConfig() (err error) {
	// Initialises config
	var configModel *config.Model
	err = utils.ReadJSONFile("config/tier/"+strings.ToLower(os.Getenv("TIER")+".json"), &configModel)
	if err != nil {
		return err
	}

	config.Set(configModel)

	// Returns
	return nil
}

func startServers() {
	// Starts HTTP server
	router.BuildRouter()
}
