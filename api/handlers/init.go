package handlers

import (
	"project/storage"
	"project/storage/managers"
)

type HTTPHandler struct {
	Chat  *managers.ChatManager
	Match *managers.MatchManager
}

func NewHandler(dbs *storage.Storage) *HTTPHandler {
	return &HTTPHandler{Chat: dbs.Chat, Match: dbs.Match}
}
