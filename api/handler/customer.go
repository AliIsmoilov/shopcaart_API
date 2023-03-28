package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Customer godoc
// @ID create_customer
// @Router /customer [POST]
// @Summary Create customer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.CreateCustomer true "CreateCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCustomer(c *gin.Context) {

	var createCustomer models.CreateCustomer

	err := c.ShouldBindJSON(&createCustomer)
	if err != nil{
		h.handlerResponse(c, "Create Customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Customer().CreateCustomer(context.Background(), &createCustomer)
	if err != nil{
		h.handlerResponse(c, "Storage Crate Customer", 500, err.Error())
		return
	}

	
	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Create Customer GET_BY_ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Customer", http.StatusCreated, resp)
}


// Get By ID Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [GET]
// @Summary Get By ID Customer
// @Description Get By ID Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCustomer(c *gin.Context) {

	id := c.Param("id")
	
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Customer Get By Id", http.StatusBadRequest, "Invali UUID")
		return
	}

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Customer Get By Id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Customer Get By Id", http.StatusOK, resp)
}


// Get List Customer godoc
// @ID get_list_customer
// @Router /customer [GET]
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCustomer(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List Customer", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List Customer", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storages.Customer().GetListCustomer(context.Background(), &models.GetListCustomerRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})
	if err != nil{
		h.handlerResponse(c, "Storage GEt List Customer", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get List Customer", http.StatusOK, resp)
}


// Update Customer godoc
// @ID update_customer
// @Router /customer/{id} [PUT]
// @Summary Update Customer
// @Description Update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.UpdateCustomer true "UpdateCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCustomer(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Update Customer", http.StatusBadRequest, "Invalid UUID")
		return
	}

	var updatecustomer models.UpdateCustomer

	err := c.ShouldBindJSON(&updatecustomer)
	if err != nil{
		h.handlerResponse(c, "Update Customer", http.StatusBadRequest, err.Error())
		return
	}

	updatecustomer.Id = id

	rowsAffected, err := h.storages.Customer().UpdateCustomer(context.Background(), &updatecustomer)
	if err != nil{
		h.handlerResponse(c, "Storage Update Customer", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Customer", http.StatusBadRequest, "No rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get By Id Update Customer", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Update Customer", http.StatusAccepted, resp)
}


// Delete Customer godoc
// @ID delete_customer
// @Router /customer/{id} [DELETE]
// @Summary Delete Customer
// @Description Delete Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Delete Customer", http.StatusBadRequest, "Invalid UUID")
		return 
	}

	err := h.storages.Customer().DeleteCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Customer", http.StatusServiceUnavailable, err.Error())
		return
	}

	h.handlerResponse(c, "Delete Customer", http.StatusOK, nil)
}