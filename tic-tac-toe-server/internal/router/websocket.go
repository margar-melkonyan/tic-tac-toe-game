package router

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Manager struct {
}

func NewManger() *Manager {
	return &Manager{}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	slog.Info("new connection from", r.RemoteAddr)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	conn.Close()
}
