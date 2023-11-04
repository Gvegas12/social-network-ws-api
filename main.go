package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Структура сервера
type Server struct {
	conns map[*websocket.Conn]bool
}

// Создание сервера
func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

// Обработчик сообщений ws
func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new message from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

// Чтение и обработка сообщений
func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error:", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()

	fmt.Println("server started")
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(":8080", nil)
}
