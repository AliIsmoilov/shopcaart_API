package handler

import (
	"net/http"
	"context"
	
	"app/api/models"
	"app/pkg/helper"

	"github.com/gin-gonic/gin"
)

// Create Author godoc
// @ID create_author
// @Router /author [POST]
// @Summary Create Author
// @Description Create Author
// @Tags Author
// @Accept json
// @Produce json
// @Param author body models.CreateAuthor true "CreateAuthorRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAuhtor(c *gin.Context) {

	var createAuhor models.CreateAuthor

	err := c.ShouldBindJSON(&createAuhor)
	if err != nil{
		h.handlerResponse(c, "Create Auhtor", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Author().CreateAuthor(context.Background(), &createAuhor)
	if err != nil{
		h.handlerResponse(c, "Create Auhtor Storage", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Author().AuthorGetById(context.Background(), &models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create Author Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Author", http.StatusCreated, resp)

}


// Get List Author godoc
// @ID get_list_author
// @Router /author [GET]
// @Summary Get List Author
// @Description Get List Author
// @Tags Author
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListAuthor(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Param("offset"))
	if err != nil{
		h.handlerResponse(c, "Get List authot Offset", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := h.getLimitQuery(c.Param("limit"))
	if err != nil{
		h.handlerResponse(c, "Get List author Limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storages.Author().GetListAuthor(context.Background(), &models.GetListAuthorRequest{
		Offset: offset,
		Limit: limit,
		Search: c.Param("search"),
	})

	if err != nil{
		h.handlerResponse(c, "Get List Author Storage", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get List Author", http.StatusOK, resp)
}


// Get By ID Author godoc
// @ID get_by_id_author
// @Router /author/{id} [GET]
// @Summary Get By ID Author
// @Description Get By ID Author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AuthorGetById(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Auhtor Get By Id", http.StatusBadRequest, "Invalid UUID")
		return
	}

	resp, err := h.storages.Author().AuthorGetById(context.Background(), &models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get Author By id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get Author By id", http.StatusOK, resp)
}



// Update Author godoc
// @ID update_author
// @Router /author/{id} [PUT]
// @Summary Update Author
// @Description Update Author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.UpdateAuthor true "UpdateAuthorkRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateAuthor(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Update Author", http.StatusBadRequest, "Invalid Author Id")
		return
	}

	var updateAuthor models.UpdateAuthor

	err := c.ShouldBindJSON(&updateAuthor)
	if err != nil{
		h.handlerResponse(c, "Update Author", http.StatusBadRequest, err.Error())
	}

	updateAuthor.Id = id

	rowsAffected, err := h.storages.Author().UpdateAuthor(context.Background(), &updateAuthor)
	if err != nil{
		h.handlerResponse(c, "Update Author Storage", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Author", http.StatusBadRequest, "No Rows Affected")
		return
	}

	resp, err := h.storages.Author().AuthorGetById(context.Background(), &models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Update Author Get By ID", http.StatusInternalServerError, err.Error())
		return
	}	

	h.handlerResponse(c, "Update Author", http.StatusAccepted, resp)

}


// Delete Author godoc
// @ID delete_author
// @Router /author/{id} [DELETE]
// @Summary Delete Author
// @Description Delete Author
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error
func (h *Handler) DeleteAuthor(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "Delete Author", http.StatusBadRequest, "Invali UUID")
		return
	}

	err := h.storages.Author().DeleteAuthor(context.Background(), &models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Delete Author", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Delete Author", http.StatusOK, nil)

}