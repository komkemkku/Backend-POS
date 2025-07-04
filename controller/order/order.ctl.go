package order

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func ListOrders(c *gin.Context) {
	var req requests.OrderRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, total, err := ListOrderService(c.Request.Context(), req)
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

func GetOrderById(c *gin.Context) {
	var req requests.OrderIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := GetOrderByIdService(c.Request.Context(), req.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateOrder(c *gin.Context) {
	staffID := c.GetInt("staff_id")
	var req requests.OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := CreateOrderService(c.Request.Context(), staffID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

// PublicCreateOrder สำหรับลูกค้าสร้างออเดอร์ (public - ไม่ต้อง auth)
func PublicCreateOrder(c *gin.Context) {
	var req requests.PublicOrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := PublicCreateOrderService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateOrder(c *gin.Context) {
	staffID := c.GetInt("staff_id")
	var req requests.OrderUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := UpdateOrderService(c.Request.Context(), req.ID, staffID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteOrder(c *gin.Context) {
	var req requests.OrderDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteOrderService(c.Request.Context(), req.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Delete Success")
}
