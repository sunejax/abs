package aerospike

import (
	"errors"
	aero "github.com/aerospike/aerospike-client-go"

	"repo.abs.com/hiring/allocate/src/config"
	dalModels "repo.abs.com/hiring/allocate/src/dal/models"
)

// Client ...
type Client struct {
	Db *aero.Client
}

// DbClient ...
var DbClient = Client{}

// Initialize ...
func Initialize() error {
	err := connect()
	if err != nil {
		return errors.New("db connect failed ->" + err.Error())
	}

	return nil
}

func connect() (err error) {
	connPolicy := aero.NewClientPolicy()
	connPolicy.User = config.GConfig.DataStores.Cart.User
	connPolicy.Password = config.GConfig.DataStores.Cart.Password

	hosts := config.GConfig.DataStores.Cart.Hosts

	var aeroHosts []*aero.Host

	for _, host := range hosts {
		aeroHost := aero.NewHost(host, config.GConfig.DataStores.Cart.Port)
		aeroHosts = append(aeroHosts, aeroHost)
	}

	DbClient.Db, err = aero.NewClientWithPolicyAndHost(connPolicy, aeroHosts...)
	if err != nil {
		return err
	}

	return nil
}

//// Create ...
//func Create(Collection string, ID string, fields map[string]interface{}, options *dalModels.Options) error {
//	// Initializes policy
//	policy := aero.NewWritePolicy(0, 0)
//
//	// CREATE_ONLY means: Create only. Fail if record exists.
//	policy.RecordExistsAction = aero.CREATE_ONLY
//
//	// Forms key
//	key, err := aero.NewKey(config.GConfig.DataStores.Cart.Database, Collection, ID)
//	if err != nil {
//		return err
//	}
//
//	// Checks options
//	// TTL
//	if options.TTL != 0 {
//		policy.Expiration = uint32(options.TTL)
//	}
//
//	// Forms bins
//	var bins []*aero.Bin
//	for key, value := range fields {
//		bin := aero.NewBin(key, value)
//		bins = append(bins, bin)
//	}
//
//	// Creates
//	err = DbClient.Db.PutBins(policy, key, bins...)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// Read ...
func Read(ID string) (map[string]interface{}, *dalModels.Options, error) {
	// Initializes policy
	policy := aero.NewPolicy()

	// Forms key
	key, err := aero.NewKey(config.GConfig.DataStores.Cart.Database, config.GConfig.DataStores.Cart.Collections.Cart.Collection, ID)
	if err != nil {
		return nil, nil, err
	}

	// Gets
	record, err := DbClient.Db.Get(policy, key)
	if err != nil {
		return nil, nil, err
	}

	// Forms meta options
	dbOptions := &dalModels.Options{
		Generation: record.Generation,
	}

	return record.Bins, dbOptions, nil
}

// ExecuteModule ...
func ExecuteModule(Index string, args ...interface{}) (*interface{}, error) {
	var moduleArgs aero.ValueArray
	for _, arg := range args {
		moduleArgs = append(moduleArgs, aero.NewValue(arg))
	}

	key, err := getKey(Index)
	if err != nil {
		return nil, err
	}
	policy := aero.NewWritePolicy(0, 0)

	result, err := DbClient.Db.Execute(policy, key, config.GConfig.DataStores.Cart.Collections.Cart.ModuleName, config.GConfig.DataStores.Cart.Collections.Cart.FunctionName, moduleArgs...)

	// Returns ...
	return &result, err
}

// GetKey ...
func getKey(Index string) (*aero.Key, error) {
	// Forms key ...
	key, err := aero.NewKey(config.GConfig.DataStores.Cart.Database, config.GConfig.DataStores.Cart.Collections.Cart.Collection, Index)
	if err != nil {
		return nil, err
	}

	// Returns ...
	return key, nil
}
