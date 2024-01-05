package json

import (
	"encoding/json"
	"net/http"

	paginationHelper "github.com/humamalamin/test-case-dating/helpers/pagination"
)

func WriteJSON(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	// enc.SetEscapeHTML(false)
	return enc.Encode(data)
}

func SuccessResponse(
	w http.ResponseWriter,
	code string,
	httpStatus int,
	data interface{},
	message string,
	pagination *paginationHelper.Page,
) {
	meta := &meta{
		Status:     true,
		Code:       code,
		Message:    message,
		Pagination: pagination,
	}

	res := &response{
		Meta: *meta,
		Data: data,
	}

	WriteJSON(w, httpStatus, res)
}

func ErrorResponse(w http.ResponseWriter, code string, httpStatus int, data interface{}, message string) {
	meta := &meta{
		Status:  false,
		Code:    code,
		Message: message,
	}

	res := &response{
		Meta: *meta,
		Data: data,
	}

	WriteJSON(w, httpStatus, res)
}
