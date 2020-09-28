package getquantity

import (
	"context"
	"errors"

	"github.com/mitchellh/mapstructure"

	dalModels "repo.abs.com/hiring/getQuantity/src/dal/models"
	"repo.abs.com/hiring/getQuantity/src/dal/store"
	"repo.abs.com/hiring/getQuantity/src/models/contracts"
)

func GetQuantity(ctx context.Context, req map[string]interface{}) (response *contracts.GetQuantityResponse, err error) {
	request, err := formInput(req)
	if err != nil {
		return nil, err
	}

	product, err := getProduct(ctx, request.ProductCode)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	response = formResponse(product.Quantity)
	return response, nil
}

func getProduct(ctx context.Context, productCode string) (*dalModels.Product, error) {
	product, err := store.GetProduct(ctx, productCode)
	return product, err
}

func formInput(req map[string]interface{}) (request *contracts.GetQuantityRequest, err error) {
	err = mapstructure.Decode(req, &request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func formResponse(quantity int64) (response *contracts.GetQuantityResponse) {
	response = &contracts.GetQuantityResponse{
		Quantity: quantity,
	}
	return response
}
