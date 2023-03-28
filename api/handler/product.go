package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.CreateProduct true "CreateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context){

	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct)
	if err != nil{
		h.handlerResponse(c, "Create Product", 400, err.Error())
		return
	}

	id, err := h.storages.Product().CreateProduct(context.Background(), &createProduct)
	if err != nil{
		h.handlerResponse(c, "Storage Create Product", 500, err.Error())
		return
	}

	resp, err := h.storages.Product().GetByIdProduct(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create Product Storage Get By Id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Product", http.StatusCreated, resp)
}

// Get By ID Product godoc
// @ID get_by_id_product
// @Router /product/{id} [GET]
// @Summary Get By ID Product
// @Description Get By ID Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdProduct(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Product Get By Id", 400, "Invalid UUID")
		return
	}

	resp, err := h.storages.Product().GetByIdProduct(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Product Get By id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Product Get By Id", http.StatusOK, resp)
}


// Get List Product godoc
// @ID get_list_product
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List Product", 400, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List Product", 400, err.Error())
		return
	}

	resp, err := h.storages.Product().GetListProduct(context.Background(), &models.GetListProductRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Storage Get List Product", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get List Product", http.StatusOK, resp)
}


// Update Product godoc
// @ID update_product
// @Router /product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param product body models.UpdateProduct true "UpdateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Update Product", 400, "Invalid UUID")
		return
	}

	var update_product models.UpdateProduct

	err := c.ShouldBindJSON(&update_product)
	if err != nil{
		h.handlerResponse(c, "Update Product", 400, err.Error())
		return 
	}

	update_product.Id = id

	rowsAffected, err := h.storages.Product().UpdateProduct(context.Background(), &update_product)
	if err != nil{
		h.handlerResponse(c, "Storage Update Product", 500, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Product", 400, "No Rows Affected")
		return
	}

	resp, err := h.storages.Product().GetByIdProduct(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Update Product Get By Id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Update Product", http.StatusAccepted, resp)
}


// Delete Product godoc
// @ID delete_product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Delete Product", http.StatusBadRequest, "Invalid UUID")
		return 
	}

	err := h.storages.Product().DeleteProduct(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Product", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Delete Product", http.StatusOK, nil)
	
}