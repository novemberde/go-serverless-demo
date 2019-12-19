package api

import "go-serverless-demo/internal/db"

type API struct {
	db *db.DB
}

func NewAPI(db *db.DB) *API {
	return &API{db}
}
