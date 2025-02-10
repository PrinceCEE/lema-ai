package pagination

import (
	"errors"
	"math"
	"strconv"

	"github.com/princecee/lema-ai/internal/db/models"
)

type PaginationQuery struct {
	Page  *int
	Limit *int
}

func GetPaginationData(opts PaginationQuery) int {
	if opts.Limit == nil {
		l := 10
		opts.Limit = &l
	}

	if opts.Page == nil {
		p := 1
		opts.Page = &p
	}

	offset := (*opts.Page - 1) * (*opts.Limit)
	return offset
}

func FormatPaginationQuery(page, limit string) (int, int, error) {
	p := 1
	l := 10

	if page != "" {
		_page, err := strconv.Atoi(page)
		if err != nil {
			return 0, 0, errors.New("invalid page number")
		}
		p = _page
	}

	if limit != "" {
		_limit, err := strconv.Atoi(limit)
		if err != nil {
			return 0, 0, errors.New("invalid limit number")
		}
		l = _limit
	}

	return p, l, nil
}

func GetTotalPages(count int64, limit int) int64 {
	return int64(math.Ceil(float64(count) / float64(limit)))
}

type GetUsersResult struct {
	Users      []*models.User `json:"users"`
	Count      int64          `json:"count"`
	TotalPages int64          `json:"total_pages"`
	Page       int64          `json:"page"`
	Limit      int64          `json:"limit"`
	HasNext    bool           `json:"has_next"`
	HasPrev    bool           `json:"has_prev"`
}
