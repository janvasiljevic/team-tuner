package out

type PaginationOut struct {
	TotalCount  int         `json:"totalCount" validate:"required"`
	Offset      int         `json:"offset" validate:"required"`
	Limit       int         `json:"limit" validate:"required"`
	CurrentPage int         `json:"currentPage" validate:"required"`
	TotalPages  int         `json:"totalPages" validate:"required"`
	Content     interface{} `json:"content" validate:"required"`
}

func NewPaginationOut(totalElementCount int, page int, pageSize int, content interface{}) *PaginationOut {
	return &PaginationOut{
		CurrentPage: page,
		TotalPages:  (totalElementCount + pageSize - 1) / pageSize, // Ensure correct calculation when not divisible evenly
		TotalCount:  totalElementCount,
		Offset:      (page) * pageSize,
		Limit:       pageSize,
		Content:     content,
	}
}
