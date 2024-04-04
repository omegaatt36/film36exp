package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/omegaatt36/film36exp/app/api/film"
	"github.com/omegaatt36/film36exp/app/api/user"
	"github.com/omegaatt36/film36exp/domain/stub"
	"github.com/omegaatt36/film36exp/logging"
	filmService "github.com/omegaatt36/film36exp/service/film"
	userService "github.com/omegaatt36/film36exp/service/user"

	"github.com/gin-gonic/gin"
)

// Server is an api server
type Server struct {
	router *gin.Engine

	filmController *film.Controller
	userController *user.Controller
}

// NewServer creates a new server
func NewServer() *Server {
	apiEngine := gin.New()
	apiEngine.RedirectTrailingSlash = true

	userRepo := stub.NewInMemoryUserRepository()

	// FIXME implement real repo
	filmController := film.NewController(filmService.NewService(
		userRepo,
		stub.NewInMemoryFilmRepository(),
	))
	userController := user.NewController(userService.NewService(
		userRepo,
	))

	return &Server{
		router: apiEngine,

		filmController: filmController,
		userController: userController,
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
	v1 := s.router.Group("v1")

	groupFilmLog := v1.Group("/film_log")
	groupFilmLog.POST("/", s.filmController.CreateFilmLog)
	groupFilmLog.GET("/:id", s.filmController.GetFilmLog)
	groupFilmLog.PATCH("/:id", s.filmController.UpdateFilmLog)
	groupFilmLog.DELETE("/:id", s.filmController.DeleteFilmLog)

	groupUser := v1.Group("/user")
	groupUser.POST("/", s.userController.CreateUser)
	groupUser.GET("/:id", s.userController.GetUser)
	groupUser.PATCH("/:id", s.userController.UpdateUser)
	groupUser.DELETE("/:id", s.userController.DeleteUser)
}
