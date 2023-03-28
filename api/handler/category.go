package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Category godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "CreateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory)
	if err != nil{
		h.handlerResponse(c, "Create Category", 400, err.Error())
		return
	}

	id, err := h.storages.Category().CreateCategory(context.Background(), &createCategory)
	if err != nil{
		h.handlerResponse(c, "Storage Create Category", 500, err.Error())
		return
	}

	resp, err := h.storages.Category().GetByIdCategory(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Create Category Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Category", http.StatusCreated, resp)
}


// Get By ID Category godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Get By ID Category", 400, "Invalid UUID")
		return
	}

	resp, err := h.storages.Category().GetByIdCategory(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Category Get By ID", http.StatusOK, resp)
}


// Get List Category godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List Category", 400, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List Category", 400, err.Error())
		return
	}

	resp, err := h.storages.Category().GetListCategory(context.Background(), &models.GetListCatogoryRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Storage Get List Category", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get List Category", http.StatusOK, resp)
}


// Update Category godoc
// @ID update_category
// @Router /category/{id} [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.UpdateCategory true "UpdateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCategory(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Update Category", 400, "Invalid UUID")
		return
	}

	var update_category models.UpdateCategory

	err := c.ShouldBindJSON(&update_category)
	if err != nil{
		h.handlerResponse(c, "Update category", 400, err.Error())
		return
	}

	update_category.Id = id

	rowsAffected, err := h.storages.Category().UpdateCategory(context.Background(), &update_category)
	if err != nil{
		h.handlerResponse(c, "Storage Update Category", 500, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Category", 400, "No rows Affected")
		return
	}

	resp, err := h.storages.Category().GetByIdCategory(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Update Category Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Update Category", http.StatusAccepted, resp)
}


// Delete Category godoc
// @ID delete_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")
	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Delete Category", 400, "Invalid UUID")
		return
	}

	err := h.storages.Category().DeleteCategory(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Category", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Delete Category", http.StatusOK, nil)
}