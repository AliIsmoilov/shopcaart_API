package api

import (
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/book", handler.CreateBook)
	r.GET("/book/:id", handler.GetByIdBook)
	r.GET("/book", handler.GetListBook)
	r.PUT("/book/:id", handler.UpdateBook)
	r.DELETE("/book/:id", handler.DeleteBook)

}

func NewApiUser(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/user", handler.CreateUser)
	r.GET("/user", handler.GetListUSer)
	r.GET("/user/:id", handler.GetByIDUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
}

func NewApiAuthor(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/author", handler.CreateAuhtor)
	r.GET("/author", handler.GetListAuthor)
	r.GET("/author/:id", handler.AuthorGetById)
	r.PUT("/author/:id", handler.UpdateAuthor)
}
