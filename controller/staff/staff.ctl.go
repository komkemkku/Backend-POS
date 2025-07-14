package staff

import (
	"Backend-POS/model"
	"Backend-POS/requests"
	response "Backend-POS/responses"

	"github.com/gin-gonic/gin"
)

func GetInfoStaff(c *gin.Context) {
	staff := c.GetInt("staff_id")

	data, err := GetByIdStaffService(c, staff)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}

func StaffList(c *gin.Context) {
	req := requests.StaffRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListStaffService(c.Request.Context(), req)
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

func GetStaffByID(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdStaffService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func CreateStaff(c *gin.Context) {
	req := requests.StaffCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateStaffService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "staff created successfully")
}

func UpdateStaff(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.StaffUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateStaffService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Updated successfully")
}

func DeleteStaff(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteStaffService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}
