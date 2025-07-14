package requests

type CategoryRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CategoryIdRequest struct {
	ID int `uri:"id"`
}

type CategoryCreateRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	DisplayOrder string `json:"display_order"`
}

type CategoryUpdateRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DisplayOrder string `json:"display_order"`
}
