package storage

import (
	"app/api/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	User() UserRepoI
	Author() AuthorRepoI
}

type BookRepoI interface {
	Create(*models.CreateBook) (string, error)
	GetByID(*models.BookPrimaryKey) (*models.Book, error)
	GetList(*models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(*models.UpdateBook) (int64, error)
	Delete(*models.BookPrimaryKey) error
}

type UserRepoI interface{
	CreateUser(*models.CreateUser) (string, error)
	UpdateUser(*models.UpdateUser) (int64, error)
	DeleteUser(*models.UserPrimaryKey) error
	UserGetByID(*models.UserPrimaryKey) (*models.User, error)
	UserGetList(*models.GetListUserRequest) (*models.GetListUserResponse, error)
}

type AuthorRepoI interface {
	CreateAuthor(*models.CreateAuthor) (string, error)
	AuthorGetById(*models.AuthorPrimaryKey) (*models.Author, error)
	GetListAuthor(*models.GetListAuthorRequest) (*models.GetListAuthorResponse, error)
	UpdateAuthor(*models.UpdateAuthor) (int64, error)
	
}
