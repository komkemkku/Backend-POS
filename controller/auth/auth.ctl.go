package auth

import (
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"Backend-POS/utils/jwt"

	"github.com/gin-gonic/gin"
)

func LoginStaff(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "ข้อมูลไม่ถูกต้อง: "+err.Error())
		return
	}

	data, err := LoginUserService(c, req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenStaff(c, data)
	if err != nil {
		response.InternalError(c, "ไม่สามารถสร้าง token ได้")
		return
	}

	// Response ตาม format ที่กำหนด
	responseData := gin.H{
		"token": token,
		"staff": gin.H{
			"id":        data.ID,
			"username":  data.UserName,
			"full_name": data.FullName,
			"role":      data.Role,
		},
	}

	response.Success(c, responseData)
}
