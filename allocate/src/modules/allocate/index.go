package allocate

import (
	"context"
	dalStore "repo.abs.com/hiring/allocate/src/dal/store"

	"github.com/mitchellh/mapstructure"

	dalCart "repo.abs.com/hiring/allocate/src/dal/cart"
	"repo.abs.com/hiring/allocate/src/models/contracts"
)

func Allocate(ctx context.Context, req map[string]interface{}) (response *contracts.AllocateResponse, err error) {
	request, err := formInput(req)
	if err != nil {
		return nil, err
	}

	err = doAllocation(ctx, request.ProductCode)
	if err != nil {
		return nil, err
	}

	err = insertInUserRecord(ctx, request.User, request.ProductCode)
	if err != nil {
		return nil, err
	}

	response = formResponse()
	return response, nil
}

func doAllocation(ctx context.Context, productCode string) error {
	err := dalCart.AllocateCart(ctx, productCode)
	return err
}

func insertInUserRecord(ctx context.Context, user, productCode string) error {
	err := dalStore.AddPurchase(ctx, user, productCode)
	return err
}

func formInput(req map[string]interface{}) (request *contracts.AllocateRequest, err error) {
	err = mapstructure.Decode(req, &request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func formResponse() (response *contracts.AllocateResponse) {
	response = &contracts.AllocateResponse{
		Allocated: true,
	}
	return response
}
