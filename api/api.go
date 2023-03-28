package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/book", handler.CreateBook)
	r.GET("/book/:id", handler.GetByIdBook)
	r.GET("/book", handler.GetListBook)
	r.PUT("/book/:id", handler.UpdateBook)
	r.DELETE("/book/:id", handler.DeleteBook)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
	r.DELETE("/author/:id", handler.DeleteAuthor)
}

func NewApiCustomer(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)
}

func NewApiCourier(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/courier", handler.CreateCourier)
	r.GET("/courier/:id", handler.GetByIDCourier)
	r.GET("/courier", handler.GetListCourier)
	r.PUT("/courier/:id", handler.UpdateCourier)
	r.DELETE("/courier/:id", handler.DeleteCourier)
}

func NewApiProduct(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)
}

func NewApiCategory(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)
}

func NewApiOrder(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrders)
	r.PUT("/order/:id", handler.UpdateOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)
	r.PATCH("/order/:id", handler.UpdatePatchOrder)
}