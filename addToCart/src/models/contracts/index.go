package contracts

// AddToCartRequest...
type AddToCartRequest struct {
	User        string
	ProductCode string
}

// AddToCartResponse...
type AddToCartResponse struct {
	ProductCode string
}
