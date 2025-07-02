package table

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func ListTables(c *gin.Context) {
	var req requests.TableRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, total, err := ListTableService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	paginate := model.Paginate{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}

	response.SuccessWithPaginate(c, data, paginate)
}

func GetTableById(c *gin.Context) {
	var req requests.TableIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := GetTableByIdService(c.Request.Context(), req.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateTable(c *gin.Context) {
	var req requests.TableCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := CreateTableService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateTable(c *gin.Context) {
	var req requests.TableUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := UpdateTableService(c.Request.Context(), req.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteTable(c *gin.Context) {
	var req requests.TableDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteTableService(c.Request.Context(), req.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Delete Success")
}

// PublicMenuByQrCode สำหรับลูกค้าที่สแกน QR Code โต๊ะ (public)
func PublicMenuByQrCode(c *gin.Context) {
	qrCode := c.Param("qrCodeIdentifier")
	data, err := PublicGetMenuByQrCodeService(c.Request.Context(), qrCode)
	if err != nil {
		response.BadRequest(c, "ไม่พบโต๊ะหรือเมนู")
		return
	}
	response.Success(c, data)
}
