package server

import (
	"github.com/leandersteiner/iot-backend/internal/model"
	"github.com/leandersteiner/iot-backend/internal/repository"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	server       *http.Server
	heartRateLog repository.HeartRateLog
}

func NewServer() *Server {
	return &Server{
		heartRateLog: repository.HeartRateLog{
			Mut:        sync.RWMutex{},
			HeartRates: []model.HeartRate{},
		},
	}
}

func (s *Server) Run(port string) error {
	router := newRouter()

	s.server = &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  90 * time.Second,
		Handler:      router,
	}

	return s.server.ListenAndServe()
}
