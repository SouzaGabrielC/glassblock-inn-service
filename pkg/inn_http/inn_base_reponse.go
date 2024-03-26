package inn_http

type InnBaseResponse struct {
	Act             string `json:"act"`
	ProtocolVersion string `json:"protocolVersion"`
	Testnet         bool   `json:"testnet"`
	Version         string `json:"version"`
}

type InnBasePaginatedResponse struct {
	Pagination InnPagination
}
