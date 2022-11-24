package pagination

import "github.com/gin-gonic/gin"

// PageRequest represents pagination request parameters.
type PageRequest struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

// NewPageRequestFromURL populates, validates and returns a page request struct.
// It extracts the data from either the query (if GET is used) or from JSON/XML
// or Form-Data (if POST is used). In case the sent data is invalid, it populates
// it with the default data.
func NewPageRequestFromURL(c *gin.Context) PageRequest {
	var pr PageRequest

	if err := c.ShouldBind(&pr); err != nil {
		pr.Page = 1
		pr.PerPage = DefaultPageSize
	}

	if pr.Page < 1 {
		pr.Page = 1
	}

	if pr.PerPage < 1 {
		pr.PerPage = DefaultPageSize
	} else if pr.PerPage > MaxPageSize {
		pr.PerPage = MaxPageSize
	}

	return pr
}
