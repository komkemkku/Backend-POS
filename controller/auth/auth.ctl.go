package auth

import (
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"Backend-POS/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginStaff(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {

		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginUserService(c, req)
	if err != nil {

		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenStaff(c, data)
	if err != nil {

		response.InternalError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}
