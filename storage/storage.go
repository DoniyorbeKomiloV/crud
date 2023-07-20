package storage

import "user/models"

type StorageI interface {
	Close()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.UserPrimaryKey, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.GetListRequest) (*models.GetListResponse, error)
	Update(*models.User) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}
