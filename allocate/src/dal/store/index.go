package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"repo.abs.com/hiring/allocate/src/config"
	"repo.abs.com/hiring/allocate/src/dal"
)

// AddPurchase ...
func AddPurchase(ctx context.Context, user string, productCode string) error {
	// Forms query
	filter := map[string]interface{}{
		"user": user,
	}
	update := bson.D{primitive.E{
		Key: "$push", Value:
		primitive.M{
			"purchase": bson.D{
				primitive.E{
					Key:   "productCode",
					Value: productCode,
				},
			},
		},
	}}

	// Hits DB
	_, err := dal.DbClient.Db.Collection(config.GConfig.DataStores.Store.Collections.User).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
