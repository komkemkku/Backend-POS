package table

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"strconv"

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
	idParam := c.Param("id")
	if idParam == "" && req.ID == 0 {
		response.BadRequest(c, "ต้องระบุ id")
		return
	}
	// แปลง id จาก param เป็น int ถ้ามี
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.BadRequest(c, "id ไม่ถูกต้อง")
			return
		}
		req.ID = id
	}
	data, err := UpdateTableService(c.Request.Context(), req.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteTable(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		response.BadRequest(c, "ต้องระบุ id")
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "id ไม่ถูกต้อง")
		return
	}
	err = DeleteTableService(c.Request.Context(), id)
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
