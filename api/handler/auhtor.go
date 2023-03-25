package handler

import (
	"app/api/models"
	"app/pkg/helper"
	// "app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAuhtor(c *gin.Context) {

	var createAuhor models.CreateAuthor

	err := c.ShouldBindJSON(&createAuhor)
	if err != nil{
		h.handlerResponse(c, "Create Auhtor", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Author().CreateAuthor(&createAuhor)
	if err != nil{
		h.handlerResponse(c, "Create Auhtor Storage", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Author().AuthorGetById(&models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Create Author Get By ID", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Create Author", http.StatusCreated, resp)

}

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

	resp, err := h.storages.Author().GetListAuthor(&models.GetListAuthorRequest{
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

func (h *Handler) AuthorGetById(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id){
		h.handlerResponse(c, "Auhtor Get By Id", http.StatusBadRequest, "Invalid UUID")
		return
	}

	resp, err := h.storages.Author().AuthorGetById(&models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Get Author By id", 500, err.Error())
		return
	}

	h.handlerResponse(c, "Get Author By id", http.StatusOK, resp)
}


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

	rowsAffected, err := h.storages.Author().UpdateAuthor(&updateAuthor)
	if err != nil{
		h.handlerResponse(c, "Update Author Storage", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0{
		h.handlerResponse(c, "Update Author", http.StatusBadRequest, "No Rows Affected")
		return
	}

	resp, err := h.storages.Author().AuthorGetById(&models.AuthorPrimaryKey{Id: id})
	if err != nil{
		h.handlerResponse(c, "Storage Update Author Get By ID", http.StatusInternalServerError, err.Error())
		return
	}	

	h.handlerResponse(c, "Update Author", http.StatusAccepted, resp)

}