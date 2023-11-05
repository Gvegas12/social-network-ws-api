package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	conns  map[*websocket.Conn]bool
}

func NewServer(config *Config) *APIServer {
	return &APIServer{
		config,
		logrus.New(),
		make(map[*websocket.Conn]bool),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server at http://localhost:" + s.config.BindAddr)
	http.Handle("/ws", websocket.Handler(s.handleWS))
	return http.ListenAndServe(s.config.BindAddr, nil)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Обработчик сообщений ws
func (s *APIServer) handleWS(ws *websocket.Conn) {
	fmt.Println("new message from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

// Чтение и обработка сообщений
func (s *APIServer) readLoop(ws *websocket.Conn) {
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

// Чтение соединения
func (s *APIServer) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error:", err)
			}
		}(ws)
	}
}
