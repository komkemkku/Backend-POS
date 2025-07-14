package menuitem

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListMenuItems(c *gin.Context) {
	var req requests.MenuItemRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, total, err := ListMenuItemService(c.Request.Context(), req)
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

func GetMenuItemByID(c *gin.Context) {
	var req requests.MenuItemIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetMenuItemByIDService(c.Request.Context(), req.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateMenuItem(c *gin.Context) {
	var req requests.MenuItemCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := CreateMenuItemService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateMenuItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req requests.MenuItemUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	data, err := UpdateMenuItemService(c.Request.Context(), id, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func DeleteMenuItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	err = DeleteMenuItemService(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Delete successfully")
}

func PublicListMenuItems(c *gin.Context) {
	var req requests.MenuItemRequest
	req.Page = 1
	req.Size = 1000

	if err := c.ShouldBindQuery(&req); err != nil {
		req.Page = 1
		req.Size = 1000
	}

	data, total, err := PublicListMenuItemService(c.Request.Context(), req)
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
