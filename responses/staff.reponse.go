package response

type StaffResponses struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
