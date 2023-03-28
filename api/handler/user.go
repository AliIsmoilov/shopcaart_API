package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create user
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param book body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context){

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil{
		h.handlerResponse(c, "Create User Body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().CreateUser(context.Background(), &createUser)
	if err != nil{
		h.handlerResponse(c, "Storage create user", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.User().UserGetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create User Get By id", http.StatusInternalServerError, err.Error())
		return
	}


	h.handlerResponse(c, "Create User", http.StatusCreated, resp)
	
}


// Get List User godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListUSer(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List User", http.StatusBadRequest, "Invalid Offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List User", http.StatusBadRequest, "Invalid Limit")
		return
	}

	
	resp, err := h.storages.User().UserGetList(context.Background(), &models.GetListUserRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Query("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Get List User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get List User", http.StatusOK, resp)
}


// Get By ID User godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Get User By Id", http.StatusBadRequest, "Invalid UUID")
		return
	}

	resp, err := h.storages.User().UserGetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	
	if err != nil{
		h.handlerResponse(c, "Get User By id", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get User By Id", http.StatusOK, resp)
}


// Update User godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UpdateUser true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Update User", http.StatusBadRequest, "Invalid User Id")
		return
	}

	var updateUser models.UpdateUser

	err := c.ShouldBindJSON(&updateUser)
	if err != nil{
		h.handlerResponse(c, "Update User", http.StatusBadRequest, err.Error())
	}

	updateUser.Id = id

	rowsAffected, err := h.storages.User().UpdateUser(context.Background(), &updateUser)
	if err != nil{
		h.handlerResponse(c, "Update User", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update User", http.StatusBadRequest, "No Rows Affected")
		return
	}

	resp, err := h.storages.User().UserGetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Update User Get By ID", http.StatusInternalServerError, err.Error())
		return
	}	

	h.handlerResponse(c, "Update User", http.StatusAccepted, resp)

}



// Delete User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Delete User", http.StatusBadRequest, "Invalid UUID")
		return
	}

	err := h.storages.User().DeleteUser(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Delete User", http.StatusOK, nil)

}