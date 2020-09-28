package cart

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	"repo.abs.com/hiring/allocate/src/config"
	aeroORM "repo.abs.com/hiring/allocate/src/orm/aerospike"
)

type Res struct {
	Success  int
	OverFlow bool
}

type record struct {
	count    int
	overflow int
}

func AllocateCart(ctx context.Context, productCode string) error {
	ok, err := inventoryStillThere(productCode)
	if err != nil {
		return err
	}
	if ok {
		result, err := aeroORM.ExecuteModule(productCode, "count", config.GConfig.Product.Inventory)
		if err != nil {
			return err
		}
		//Lua Blocking
		var res *Res
		err = mapstructure.Decode((*result).(map[interface{}]interface{}), &res)
		if err != nil {
			return err
		}
		if res.OverFlow {
			return errors.New("overflow")
		}
		return nil
	}
	return errors.New("overflow")
}

func inventoryStillThere(productCode string) (bool, error) {
	rec, _, err := aeroORM.Read(productCode)
	if rec != nil {
		var Record record
		err = mapstructure.Decode(rec, &Record)
		if err != nil {
			return false, err
		}
		if Record.overflow == 1 {
			return false, nil
		}
		return true, nil

	}
	return true, nil
}
