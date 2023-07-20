package handler

import (
	"user/config"
	"user/storage"
)

type Handler struct {
	cfg  *config.Config
	strg storage.StorageI
}

func NewHandler(cfg *config.Config, strg storage.StorageI) *Handler {
	return &Handler{
		cfg:  cfg,
		strg: strg,
	}
}
