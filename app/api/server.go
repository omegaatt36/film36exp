package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/omegaatt36/film36exp/app/api/film"
	"github.com/omegaatt36/film36exp/app/api/user"
	"github.com/omegaatt36/film36exp/logging"
	"github.com/omegaatt36/film36exp/rdb"
	"github.com/omegaatt36/film36exp/rdb/database"
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

	repo := rdb.NewGormRepo(database.GetDB(database.Default))

	filmController := film.NewController(filmService.NewService(
		repo,
		repo,
	))
	userController := user.NewController(userService.NewService(
		repo,
	))

	return &Server{
		router: apiEngine,

		filmController: filmController,
		userController: userController,
	}
}

// Start starts the server
func (s *Server) Start(ctx context.Context) <-chan struct{} {
	s.router.Use(corsMiddleware())
	s.registerRoutes()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", defaultConfig.appPort),
		Handler: s.router,
	}

	closeChain := make(chan struct{})
	go func() {
		defer func() {
			logging.Info("api stopped")
			closeChain <- struct{}{}
			close(closeChain)
		}()

		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			logging.Fatal("Server Shutdown: ", err)
		}
	}()

	logging.Info("starts serving...")

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logging.Fatalf("listen: %s\n", err)
		}
	}()

	return closeChain
}

func (s *Server) registerRoutes() {
	v1 := s.router.Group("v1")

	v1.Use(injectLogging([]string{}), recovery())

	groupFilmLog := v1.Group("/film_logs")
	groupFilmLog.POST("", s.filmController.CreateFilmLog)
	groupFilmLog.GET(":id", s.filmController.GetFilmLog)
	groupFilmLog.PUT(":id", s.filmController.UpdateFilmLog)
	groupFilmLog.DELETE(":id", s.filmController.DeleteFilmLog)

	groupPhoto := v1.Group("/photos")
	groupPhoto.POST("", s.filmController.CreatePhoto)
	groupPhoto.GET(":id", s.filmController.GetPhoto)
	groupPhoto.PUT(":id", s.filmController.UpdatePhoto)
	groupPhoto.DELETE(":id", s.filmController.DeletePhoto)

	groupUser := v1.Group("/users")
	groupUser.POST("", s.userController.CreateUser)
	groupUser.GET(":id", s.userController.GetUser)
	groupUser.PUT(":id", s.userController.UpdateUser)
	groupUser.DELETE(":id", s.userController.DeleteUser)
}
