package store

import (
	"context"
	"github.com/mitchellh/mapstructure"

	"repo.abs.com/hiring/getQuantity/src/config"
	"repo.abs.com/hiring/getQuantity/src/dal"
	"repo.abs.com/hiring/getQuantity/src/dal/models"
)

// Get ...
func GetProduct(ctx context.Context, productCode string) (*models.Product, error) {
	// Forms query
	query := map[string]interface{}{
		"productCode": productCode,
	}

	// Hits DB
	var product map[string]interface{}
	err := dal.DbClient.Db.Collection(config.GConfig.DataStores.Store.Collections.Product).FindOne(nil, query).Decode(&product)
	if err != nil {
		// Handles no record
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}

		// Handles other errors
		return nil, err
	}

	// Decodes
	var Product *models.Product
	err = mapstructure.Decode(product, &Product)
	if err != nil {
		return nil, err
	}

	// Returns
	return Product, nil
}
