package response

type CategoryResponses struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DisplayOrder string    `json:"display_order"`
}
