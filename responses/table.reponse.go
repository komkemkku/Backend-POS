package response

type TableResponses struct {
	ID               int    `json:"id"`
	TableNumber      int    `json:"table_number"`
	Capacity         int    `json:"capacity"`
	Status           string `json:"status"`
	QrCodeIdentifier string `json:"qr_code_identifier"`
}
