package requests

type TableRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type TableIdRequest struct {
	ID int `uri:"id"`
}

type TableCreateRequest struct {
	TableNumber      int    `json:"table_number"`
	Capacity         int    `json:"capacity"`
	Status           string `json:"status"`
	QrCodeIdentifier string `json:"qr_code_identifier"`
}

type TableUpdateRequest struct {
	ID               int    `json:"id"`
	TableNumber      int    `json:"table_number"`
	Capacity         int    `json:"capacity"`
	Status           string `json:"status"`
	QrCodeIdentifier string `json:"qr_code_identifier"`
}

type TableDeleteRequest struct {
	ID int `json:"id"`
}
