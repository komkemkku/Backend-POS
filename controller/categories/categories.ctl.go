package categories

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func CategoryList(c *gin.Context) {
	req := requests.CategoryRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListCategoryService(c.Request.Context(), req)
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

func GetCategoryByID(c *gin.Context) {
	id := requests.CategoryIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdCategoryService(c.Request.Context(), id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateCategory(c *gin.Context) {
	req := requests.CategoryCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := CreateCategoryService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UpdateCategory(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.CategoryUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateCategoryService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Updated successfully")
}

func DeleteCategory(c *gin.Context) {
	id := requests.CategoryIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteCategoryService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}
