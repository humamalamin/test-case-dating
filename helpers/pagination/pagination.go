package pagination

import (
	"math"
	"strconv"

	convHelper "github.com/humamalamin/test-case-dating/helpers/conv"
	validationHelper "github.com/humamalamin/test-case-dating/helpers/validation"
)

type Pagination interface {
	AddPagination(totalData int, page string, perPage string) (*Page, error)
}

type Options struct{}

// AddPagination implements Pagination.
func (o *Options) AddPagination(totalData int, page string, perPage string) (*Page, error) {
	if page == "" {
		return nil, ErrorPageEmpty
	}

	err := validationHelper.StrIsDigit(page)
	if err != nil {
		return nil, ErrorPageInvalid
	}

	newPage, _ := convHelper.StringToInt(page)

	if newPage <= 0 {
		return nil, ErrorPage
	}

	limitData := 10
	if perPage != "" {
		limitData, _ = strconv.Atoi(perPage)
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limitData)))

	last := (newPage * limitData)
	first := last - limitData

	if totalData < last {
		last = totalData
	}

	zeroPage := &Page{PageCount: 1, Page: newPage}
	if totalPage == 0 && newPage == 1 {
		return zeroPage, nil
	}

	if newPage > totalPage {
		return nil, ErrorMaxPage
	}

	pages := &Page{
		PageCount:  totalPage,
		Page:       newPage,
		TotalCount: totalData,
		PerPage:    limitData,
		First:      first,
		Last:       last,
	}

	return pages, nil
}

func NewPagination() Pagination {
	pagination := new(Options)

	return pagination
}
