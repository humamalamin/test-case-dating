package json

import paginationHelper "github.com/humamalamin/test-case-dating/helpers/pagination"

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status     bool                   `json:"status"`
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Pagination *paginationHelper.Page `json:"pagination,omitempty"`
}
