package pagination

var (
	// DefaultPageSize specifies the amount of entries in a single page
	DefaultPageSize = 10
	// MaxPageSize specifies the maximal amount of entries in a single page
	MaxPageSize = 1000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies query parameter name for page size
	PageSizeVar = "per_page"
)

// Page represents a paginated list of data items.
type Page[T any] struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	Items      []T `json:"items"`
}

// New creates a new Page instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// And the total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func NewPage[T any](page, perPage, total int, items []T) *Page[T] {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}

	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Page[T]{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
		Items:      items,
	}
}
