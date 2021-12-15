package paginator

import (
	"math"
	"net/http"
	"strconv"
)

const (
	defaultLimit = 10
	maxLimit     = 1000
)

type Pagination struct {
	Limit        int   `json:"limit,omitempty"`
	Offset       int   `json:"-"`
	Page         int   `json:"page,omitempty"`
	NextPage     int   `json:"next_page,omitempty"`
	PreviousPage int   `json:"previous_page,omitempty"`
	Count        int64 `json:"count,omitempty"`
	TotalPage    int   `json:"total_page,omitempty"`
}

func Paging(p *Pagination) *Pagination {
	p.TotalPage = int(math.Ceil(float64(p.Count) / float64(p.Limit)))
	if p.Page > 1 {
		p.PreviousPage = p.Page - 1
	} else {
		p.PreviousPage = p.Page
	}
	if p.Page == p.TotalPage {
		p.NextPage = p.Page
	} else {
		p.NextPage = p.Page + 1
	}
	return p
}

func Paginate(r *http.Request) (*Pagination, error) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = defaultLimit
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	p := &Pagination{
		Limit: limit,
		Page:  page,
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
	if p.Limit > 1000 {
		p.Limit = maxLimit
	}
	p.Offset = (p.Page - 1) * p.Limit
	return p, nil
}
