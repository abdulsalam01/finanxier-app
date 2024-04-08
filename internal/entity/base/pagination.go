package base

type PaginationRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PaginationResponse[T any] struct {
	Response[T]
	Total int `json:"total"`
}
