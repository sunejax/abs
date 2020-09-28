package dal

import (
	"crypto/tls"
	"errors"

	"context"
	"time"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"repo.abs.com/hiring/allocate/src/config"
)

// Client ...
type Client struct {
	Db *mongo.Database
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

func connect() error {
	mongoConnOptions := &options.ClientOptions{
		Hosts:   config.GConfig.DataStores.Store.Hosts,
		Monitor: apmmongo.CommandMonitor(),
	}

	// Checks TLS
	if config.GConfig.DataStores.Store.UseTLS {
		tlsConf := &tls.Config{
			InsecureSkipVerify: true,
		}

		mongoConnOptions.TLSConfig = tlsConf
	}

	// Checks auth
	if config.GConfig.DataStores.Store.Auth {
		creds := &options.Credential{
			Username:   config.GConfig.DataStores.Store.User,
			Password:   config.GConfig.DataStores.Store.Password,
			AuthSource: config.GConfig.DataStores.Store.AuthSource,
		}

		mongoConnOptions.Auth = creds
	}

	client, err := mongo.NewClient(mongoConnOptions)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return errors.New("client connection failed ->" + err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.New("client ping failed ->" + err.Error())
	}

	DbClient.Db = client.Database(config.GConfig.DataStores.Store.Database)

	return nil
}
