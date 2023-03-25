package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateUser(c *gin.Context){

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil{
		h.handlerResponse(c, "Create User Body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().CreateUser(&createUser)
	if err != nil{
		h.handlerResponse(c, "Storage create user", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.User().UserGetByID(&models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create User Get By id", http.StatusInternalServerError, err.Error())
		return
	}


	h.handlerResponse(c, "Create User", http.StatusCreated, resp)
	
}

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

	
	resp, err := h.storages.User().UserGetList(&models.GetListUserRequest{
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

func (h *Handler) GetByIDUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Get User By Id", http.StatusBadRequest, "Invalid UUID")
		return
	}

	resp, err := h.storages.User().UserGetByID(&models.UserPrimaryKey{Id: id})
	
	if err != nil{
		h.handlerResponse(c, "Get User By id", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get User By Id", http.StatusOK, resp)
}


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

	rowsAffected, err := h.storages.User().UpdateUser(&updateUser)
	if err != nil{
		h.handlerResponse(c, "Update User", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update User", http.StatusBadRequest, "No Rows Affected")
		return
	}

	resp, err := h.storages.User().UserGetByID(&models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Update User Get By ID", http.StatusInternalServerError, err.Error())
		return
	}	

	h.handlerResponse(c, "Update User", http.StatusAccepted, resp)

}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Delete User", http.StatusBadRequest, "Invalid UUID")
		return
	}

	err := h.storages.User().DeleteUser(&models.UserPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Delete User", http.StatusOK, nil)

}