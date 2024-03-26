package inn_http

type InnPagination struct {
	TotalCount int64 `json:"totalCount"`
	PerPage    int64 `json:"perPage"`
	Current    int64 `json:"current"`
	Last       int64 `json:"last"`
	First      int64 `json:"first"`
	Prev       int64 `json:"prev"`
}

type InnPaginationRequest struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"perPage"`
}
