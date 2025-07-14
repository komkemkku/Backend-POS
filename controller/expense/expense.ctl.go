package expense

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func ExpenseList(c *gin.Context) {
	req := requests.ExpenseRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListExpenseService(c.Request.Context(), req)
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

func GetExpenseByID(c *gin.Context) {
	id := requests.ExpenseIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdExpenseService(c.Request.Context(), id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateExpense(c *gin.Context) {
	staff := c.GetInt("staff_id")
	req := requests.ExpenseCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.StaffID = staff

	data, err := CreateExpenseService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateExpense(c *gin.Context) {
	staff := c.GetInt("staff_id")

	var uri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	var req requests.ExpenseUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.StaffID = staff

	data, err := UpdateExpenseService(c, uri.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteExpense(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := DeleteExpenseService(c.Request.Context(), uri.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Expense deleted successfully")
}
