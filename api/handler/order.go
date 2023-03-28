package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "CreateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder)
	if err != nil{
		h.handlerResponse(c, "Create Order", 400, err.Error())
	}

	id, err := h.storages.Order().CreateOrder(context.Background(), &createOrder)
	if err != nil{
		h.handlerResponse(c, "Storage Create Order", 500, err.Error())
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Create Order Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Order", http.StatusCreated, resp)
} 


// Get By ID Order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdOrder(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Order Get By Id", 400, "Invalid UUID")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get By Id Order", 500, err.Error())
	}

	h.handlerResponse(c, "Order Get By Id", http.StatusOK, resp)
}


// Get List Order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListOrders(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil {
		h.handlerResponse(c, "Get Lsit Orders", 400, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List Orders", 400, err.Error())
		return
	}

	resp, err := h.storages.Order().GetListOrders(context.Background(), &models.GetListOrderRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Storage Get List", 500, err.Error())
		return
	}
	
	h.handlerResponse(c, "Get List Order", 200, resp)
}


// Update Order godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update order
// @Description Update order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.UpdateOrder true "UpdateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Update Order", 400, "Invalid UUID")
		return
	}

	var updateOrder models.UpdateOrder

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil{
		h.handlerResponse(c, "Update Order", 400, err.Error())
		return
	}

	updateOrder.Id = id

	rowsAffected, err := h.storages.Order().UpdateOrder(context.Background(), &updateOrder)
	if err != nil{
		h.handlerResponse(c, "Storage Update Order", 500, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Order", 400, "No Rows Affected")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Update Order Get By ID", 500, err.Error())
	}

	h.handlerResponse(c, "Update Order", 200, resp)
}


// Delete Order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Delete Order", 400, "Invalid UUID")
		return
	}

	err := h.storages.Order().DeleteOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Order", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Delete Order", 200, nil)
}


// Update Patch Order godoc
// @ID update_patch_order
// @Router /order/{id} [PATCH]
// @Summary Update Patch Order
// @Description Update Patch Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.PatchRequest true "UpdatPatchOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Update Patch Order", 400, "Invalid UUID")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil{
		h.handlerResponse(c, "Update Patch Order", 400, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Order().PatchOrder(context.Background(), &object)
	if err != nil{
		h.handlerResponse(c, "Storage Patch Order", 500, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Patch Order", 400, "No rows affected")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Patch Order Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Patch Order", 200, resp)
}