package sockets

import (
	e "github.com/jsbento/chess-server/cmd/engine"
	eT "github.com/jsbento/chess-server/cmd/engine/types"
)

type ChessHub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Engine     *e.Engine
	SearchInfo *eT.SearchInfo
}

func NewChessHub() *ChessHub {
	return &ChessHub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Engine:     e.NewEngine(),
		SearchInfo: &eT.SearchInfo{},
	}
}

func (h *ChessHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
