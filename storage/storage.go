package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Book() 		BookRepoI
	User() 		UserRepoI
	Author() 	AuthorRepoI
	Customer() 	CustomerRepoI
	Courier()	CourierRepoI
	Product()	ProductRepoI
	Category()	CategoryRepoI
	Order()		OrderRepoI
}

type BookRepoI interface {
	Create(context.Context, *models.CreateBook) (string, error)
	GetByID(context.Context, *models.BookPrimaryKey) (*models.Book, error)
	GetList(context.Context, *models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(context.Context, *models.UpdateBook) (int64, error)
	Delete(context.Context, *models.BookPrimaryKey) error
}

type UserRepoI interface{
	CreateUser(context.Context, *models.CreateUser) (string, error)
	UpdateUser(context.Context, *models.UpdateUser) (int64, error)
	DeleteUser(context.Context, *models.UserPrimaryKey) error
	UserGetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	UserGetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
}

type AuthorRepoI interface {
	CreateAuthor(context.Context, *models.CreateAuthor) (string, error)
	AuthorGetById(context.Context, *models.AuthorPrimaryKey) (*models.Author, error)
	GetListAuthor(context.Context, *models.GetListAuthorRequest) (*models.GetListAuthorResponse, error)
	UpdateAuthor(context.Context, *models.UpdateAuthor) (int64, error)
	DeleteAuthor(context.Context, *models.AuthorPrimaryKey) error
	
}

type CustomerRepoI interface {
	CreateCustomer(context.Context, *models.CreateCustomer) (string, error)
	GetByIdCustomer(context.Context, *models.CustomerPrimaryKey) (*models.Customer, error)
	GetListCustomer(context.Context, *models.GetListCustomerRequest) (*models.GetListCustomerResponse, error)
	UpdateCustomer(context.Context, *models.UpdateCustomer) (int64, error)
	DeleteCustomer(context.Context, *models.CustomerPrimaryKey) (error)
}

type CourierRepoI interface {
	CreateCourier(context.Context, *models.CreateCourier) (string, error)
	GetByIDCourier(context.Context, *models.CourierPrimaryKey) (*models.Courier, error)
	GetListCourier(context.Context, *models.GetListCourierRequest) (*models.GetListCourierResponse, error)
	UpdateCourier(context.Context, *models.UpdateCourier) (int64, error)
	DeleteCourier(context.Context, *models.CourierPrimaryKey) (error)
}

type ProductRepoI interface {
	CreateProduct(context.Context, *models.CreateProduct) (string, error)
	GetByIdProduct(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetListProduct(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	UpdateProduct(context.Context, *models.UpdateProduct) (int64, error)
	DeleteProduct(context.Context, *models.ProductPrimaryKey) (error)
}

type CategoryRepoI interface {
	CreateCategory(context.Context, *models.CreateCategory) (string, error)
	GetByIdCategory(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetListCategory(context.Context, *models.GetListCatogoryRequest) (*models.GetListCategoryResponse, error)
	UpdateCategory(context.Context, *models.UpdateCategory) (int64, error)
	DeleteCategory(context.Context, *models.CategoryPrimaryKey) (error)
}

type OrderRepoI interface {
	CreateOrder(context.Context, *models.CreateOrder) (string, error)
	GetByIdOrder(context.Context, *models.OrderPrimaryKey) (*models.Order, error)
	GetListOrders(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	UpdateOrder(context.Context,*models.UpdateOrder) (int64, error)
	PatchOrder(context.Context, *models.PatchRequest) (int64, error)
	DeleteOrder(context.Context,*models.OrderPrimaryKey) (error)
}