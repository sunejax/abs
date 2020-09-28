package contracts

// GetQuantityRequest...
type GetQuantityRequest struct {
	ProductCode string
}

// GetQuantityResponse...
type GetQuantityResponse struct {
	Quantity int64
}
