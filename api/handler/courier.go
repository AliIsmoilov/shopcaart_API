package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Courier godoc
// @ID create_courier
// @Router /courier [POST]
// @Summary Create courier
// @Description Create Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param courier body models.CreateCourier true "CreateCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCourier(c *gin.Context) {

	var createCourier models.CreateCourier

	err := c.ShouldBindJSON(&createCourier)
	if err != nil{
		h.handlerResponse(c, "Create Courier", http.StatusBadRequest, err.Error())
		return 
	}

	id, err := h.storages.Courier().CreateCourier(context.Background(), &createCourier)
	if err != nil{
		h.handlerResponse(c, "Storage Create Courier", 500, err.Error())
		return
	}

	resp, err := h.storages.Courier().GetByIDCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create Courier Storage GET BY ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Courier", http.StatusCreated, resp)
}


// Get By ID Courier godoc
// @ID get_by_id_courier
// @Router /courier/{id} [GET]
// @Summary Get By ID Courier
// @Description Get By ID Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDCourier(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Get Courier By Id", http.StatusBadRequest, "Invalid UUID")
		return
	}

	resp, err := h.storages.Courier().GetByIDCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get Courier By ID", 500, err.Error())
	}
	
	h.handlerResponse(c, "Courier Get By Id", http.StatusOK, resp)
}


// Get List Courier godoc
// @ID get_list_courier
// @Router /courier [GET]
// @Summary Get List Courier
// @Description Get List Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCourier(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List Courier", 400, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List Courier", 400, err.Error())
		return
	}

	resp, err := h.storages.Courier().GetListCourier(context.Background(), &models.GetListCourierRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Storage Get List Courier", 500, err.Error())
		return 
	}

	h.handlerResponse(c, "Get List Courier", http.StatusOK, resp)
}


// Update Courier godoc
// @ID update_courier
// @Router /courier/{id} [PUT]
// @Summary Update Courier
// @Description Update Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.UpdateCourier true "UpdateCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Update Courier", http.StatusBadRequest, "Invalid UUID")
		return
	}

	var updateCourier models.UpdateCourier

	err := c.ShouldBindJSON(&updateCourier)
	if err != nil{
		h.handlerResponse(c, "ShouldBind Update Courier", http.StatusBadRequest, err.Error())
		return
	}

	updateCourier.Id = id

	rows, err := h.storages.Courier().UpdateCourier(context.Background(), &updateCourier)
	if err != nil{
		h.handlerResponse(c, "Storage Update Courier", 500, err.Error())
		return
	}

	if rows <= 0{
		h.handlerResponse(c, "Update Courier", http.StatusBadRequest, "No rows affected")
		return
	}	
	
	resp, err := h.storages.Courier().GetByIDCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Update Courier Get By Id Storage", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Update Courier", http.StatusAccepted, resp)
}


// Delete Courier godoc
// @ID delete_courier
// @Router /courier/{id} [DELETE]
// @Summary Delete Courier
// @Description Delete Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteCourier(c *gin.Context){

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Delete Courier", http.StatusBadRequest, "Invalid UUID")
		return
	}

	err := h.storages.Courier().DeleteCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Courier", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Delte Courier", 200, nil)
}