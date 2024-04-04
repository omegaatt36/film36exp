package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/omegaatt36/film36exp/logging"

	"github.com/gin-gonic/gin"
)

// Server is an api server
type Server struct {
	router *gin.Engine
}

// NewServer creates a new server
func NewServer() *Server {
	apiEngine := gin.New()
	apiEngine.RedirectTrailingSlash = true

	return &Server{
		router: apiEngine,
	}
}

// Start starts the server
func (s *Server) Start(ctx context.Context, addr string) {
	s.registerRoutes()

	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			logging.Fatal("Server Shutdown: ", err)
		}
	}()

	logging.Info("starts serving...")
	if err := srv.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		logging.Fatalf("listen: %s\n", err)
	}

}

func (s *Server) registerRoutes() {

}
