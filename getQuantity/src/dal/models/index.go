package models

// Options ...
type Options struct {
	TTL        int64
	Generation uint32
}

type Product struct {
	ProductCode string `json:"productCode" bson:"productCode" mapstructure:"productCode"`
	Quantity    int64  `json:"quantity" bson:"quantity" mapstructure:"quantity"`
}
