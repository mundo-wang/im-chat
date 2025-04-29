package utils

type PageResult[T interface{}] struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	Total     int64 `json:"total"`
	PageCount int   `json:"page_count"`
	Records   []T   `json:"records"`
}
