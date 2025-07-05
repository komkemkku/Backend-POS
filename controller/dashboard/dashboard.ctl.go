package dashboard

import (
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func GetDashboardSummary(c *gin.Context) {
	data, err := GetDashboardSummaryService(c.Request.Context())
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}
