package contracts

// AllocateRequest...
type AllocateRequest struct {
	ProductCode string
	User        string
}

// AllocateResponse...
type AllocateResponse struct {
	Allocated bool
}
