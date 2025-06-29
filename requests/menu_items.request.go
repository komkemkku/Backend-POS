package requests

type MenuItemRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type MenuItemIdRequest struct {
	ID int `uri:"id"`
}

type MenuItemCreateRequest struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"image_url"`
	IsAvailable bool   `json:"is_available"`
}

type MenuItemUpdateRequest struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"image_url"`
	IsAvailable bool   `json:"is_available"`
}

type MenuItemDeleteRequest struct {
	ID int `json:"id"`
}
