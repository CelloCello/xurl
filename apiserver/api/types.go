package api

import (
	"short_url/apiserver/pkg/database"
)

type Response struct {
	Message string      `json:"msg"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}

type LinksResponse struct {
	Response
	Payload []database.Link `json:"payload"`
}

type CreateLinkResponse struct {
	Response
	Payload database.Link `json:"payload"`
}
