package pagination

import (
	"errors"
	"strconv"
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
