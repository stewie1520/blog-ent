package types

type Pagination[T any] struct {
	Items   []T   `json:"items"`
	Total   int64 `json:"total"`
	HasMore bool  `json:"hasMore"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
}

func NewPagination[T any](items []T, total int64, offset int64, limit int64) *Pagination[T] {
	if items == nil {
		items = []T{}
	}

	return &Pagination[T]{
		Items:   items,
		Total:   total,
		HasMore: total > offset+limit,
		Offset:  offset,
		Limit:   limit,
	}
}

type PaginationParams struct {
	Offset int `form:"offset,default=0" binding:"gte=0"`
	Limit  int `form:"limit,default=20" binding:"lte=100"`
}
