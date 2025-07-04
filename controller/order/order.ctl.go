package order

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"strconv"

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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req requests.OrderUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := UpdateOrderService(c.Request.Context(), id, staffID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	err = DeleteOrderService(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Delete Success")
}

// PublicGetOrdersByTable สำหรับลูกค้าดูประวัติออเดอร์ตามโต๊ะ (public)
func PublicGetOrdersByTable(c *gin.Context) {
	qrCode := c.Param("qrCodeIdentifier")
	data, err := PublicGetOrdersByTableService(c.Request.Context(), qrCode)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, data)
}

// PublicGetOrderStatus สำหรับลูกค้าดูสถานะออเดอร์เฉพาะ (public)
func PublicGetOrderStatus(c *gin.Context) {
	orderIDParam := c.Param("orderID")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		response.BadRequest(c, "Invalid order ID")
		return
	}

	qrCode := c.Param("qrCodeIdentifier")
	data, err := PublicGetOrderStatusService(c.Request.Context(), orderID, qrCode)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, data)
}

// PublicGetAllOrderHistory สำหรับดูประวัติออเดอร์ทั้งหมด (รวมที่ชำระแล้ว)
func PublicGetAllOrderHistory(c *gin.Context) {
	qrCode := c.Param("qrCodeIdentifier")
	data, err := PublicGetAllOrderHistoryService(c.Request.Context(), qrCode)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, data)
}

// PublicClearTableHistory สำหรับล้างประวัติหลังชำระเงิน (สำหรับ staff)
func PublicClearTableHistory(c *gin.Context) {
	staffID := c.GetInt("staff_id") // จาก middleware
	qrCode := c.Param("qrCodeIdentifier")

	err := PublicClearTableHistoryService(c.Request.Context(), qrCode, staffID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "ล้างประวัติโต๊ะเรียบร้อยแล้ว")
}

// PublicGetTableSummary สำหรับดูสรุปโต๊ะ
func PublicGetTableSummary(c *gin.Context) {
	qrCode := c.Param("qrCodeIdentifier")
	data, err := PublicGetTableSummaryService(c.Request.Context(), qrCode)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, data)
}

// AdvancedClearTableHistory สำหรับล้างประวัติแบบละเอียด
func AdvancedClearTableHistory(c *gin.Context) {
	qrCode := c.Param("qrCodeIdentifier")
	clearType := c.Query("type") // payment, cancel_all, complete_all

	if clearType == "" {
		clearType = "payment" // default
	}

	staffID, exists := c.Get("staffID")
	if !exists {
		response.Unauthorized(c, "ไม่พบข้อมูล staff")
		return
	}

	data, err := AdvancedClearTableHistoryService(c.Request.Context(), qrCode, staffID.(int), clearType)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

// CancelSpecificOrder สำหรับยกเลิกออเดอร์เฉพาะ
func CancelSpecificOrder(c *gin.Context) {
	orderIDStr := c.Param("orderID")
	qrCode := c.Param("qrCodeIdentifier")
	reason := c.Query("reason")

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		response.BadRequest(c, "ID ออเดอร์ไม่ถูกต้อง")
		return
	}

	staffID, exists := c.Get("staffID")
	if !exists {
		response.Unauthorized(c, "ไม่พบข้อมูล staff")
		return
	}

	err = CancelSpecificOrderService(c.Request.Context(), orderID, qrCode, staffID.(int), reason)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "ยกเลิกออเดอร์เรียบร้อยแล้ว")
}

// UpdateOrderStatus สำหรับอัปเดตสถานะออเดอร์
func UpdateOrderStatus(c *gin.Context) {
	orderIDStr := c.Param("orderID")
	var req struct {
		Status string `json:"status" binding:"required"`
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		response.BadRequest(c, "ID ออเดอร์ไม่ถูกต้อง")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "ข้อมูลไม่ถูกต้อง")
		return
	}

	staffID, exists := c.Get("staffID")
	if !exists {
		response.Unauthorized(c, "ไม่พบข้อมูล staff")
		return
	}

	data, err := UpdateOrderStatusService(c.Request.Context(), orderID, req.Status, staffID.(int))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}
