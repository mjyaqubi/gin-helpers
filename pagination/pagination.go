package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination type struct
type Pagination struct{}

// New - Pagination
func New() *Pagination {
	return &Pagination{}
}

// Query type
type Query struct {
	Skip  int64
	Limit int64
}

// Get pagination query
func (helper *Pagination) Get(c *gin.Context) Query {
	var pageNumber int64
	var pageItems int64

	if i, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64); err == nil {
		pageNumber = i
	} else {
		pageNumber = 10
	}

	if i, err := strconv.ParseInt(c.DefaultQuery("items", "10"), 10, 64); err == nil {
		pageItems = i
	} else {
		pageItems = 10
	}

	return Query{
		Skip:  (pageNumber - 1) * pageItems,
		Limit: pageItems,
	}
}
