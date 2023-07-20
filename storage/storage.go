package storage

import "user/models"

type StorageI interface {
	Close()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.UserPrimaryKey, error)
	GetByIdUser(*models.UserPrimaryKey) (*models.User, error)
}
