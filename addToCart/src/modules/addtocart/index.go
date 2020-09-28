package addtocart

import (
	"context"
	"github.com/mitchellh/mapstructure"
	dalStore "repo.abs.com/hiring/addToCart/src/dal/store"
	"repo.abs.com/hiring/addToCart/src/models/contracts"
)

func AddToCart(ctx context.Context, req map[string]interface{}) (response *contracts.AddToCartResponse, err error) {
	request, err := formInput(req)
	if err != nil {
		return nil, err
	}

	err = pushToCart(ctx, request.User, request.ProductCode)
	if err != nil {
		return nil, err
	}
	response = formResponse(request.ProductCode)
	return response, nil
}

func pushToCart(ctx context.Context, user, productCode string) error {
	err := dalStore.AddToCart(ctx, user, productCode)
	return err
}
func formInput(req map[string]interface{}) (request *contracts.AddToCartRequest, err error) {
	err = mapstructure.Decode(req, &request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func formResponse(productCode string) (response *contracts.AddToCartResponse) {
	response = &contracts.AddToCartResponse{
		ProductCode: productCode,
	}
	return response
}
