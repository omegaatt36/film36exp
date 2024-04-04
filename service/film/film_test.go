package film_test

import (
	"context"
	"testing"
	"time"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/domain/stub"
	"github.com/omegaatt36/film36exp/service/film"
	"github.com/omegaatt36/film36exp/util"
	"github.com/stretchr/testify/suite"
)

type filmServiceTestSuite struct {
	filmService *film.Service

	suite.Suite

	userRepo domain.UserRepository
	filmRepo domain.FilmRepository
}

func (s *filmServiceTestSuite) SetupSuite() {
	s.userRepo = stub.NewInMemoryUserRepository()
	s.filmRepo = stub.NewInMemoryFilmRepository()

	s.filmService = film.NewService(
		s.userRepo,
		s.filmRepo,
	)
}

func (s *filmServiceTestSuite) TestCreateFilmWithInvalidReq() {
	s.Error(s.filmService.CreateFilmLog(context.TODO(), film.CreateFilmLogRequest{}),
		"empty req")
	s.Error(s.filmService.CreateFilmLog(context.TODO(), film.CreateFilmLogRequest{
		UserID: 9999,
	}), "without format")
	s.Error(s.filmService.CreateFilmLog(context.TODO(), film.CreateFilmLogRequest{
		UserID: 9999,
		Format: domain.FilmFormat("invalid"),
	}), "invalid format")
	s.Error(s.filmService.CreateFilmLog(context.TODO(), film.CreateFilmLogRequest{
		Format: domain.FilmFormat135,
	}), "without user id")
}

func (s *filmServiceTestSuite) TestCreateFilm() {
	req := film.CreateFilmLogRequest{
		UserID:       1,
		Format:       domain.FilmFormat135,
		Negative:     util.Pointer(true),
		Brand:        util.Pointer("Kodak"),
		ISO:          util.Pointer(400),
		PurchaseDate: util.Pointer(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		LoadDate:     util.Pointer(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		Notes:        "",
	}
	s.Error(s.filmService.CreateFilmLog(context.TODO(), req), "user not found")

	s.NoError(s.userRepo.CreateUser(context.TODO(), &domain.User{
		ID: 1,
	}))

	s.NoError(s.filmService.CreateFilmLog(context.TODO(), req))

	expectedFilmLog, err := s.filmRepo.GetFilmLog(context.TODO(), 1)
	s.NoError(err)

	s.Equal(req.UserID, expectedFilmLog.UserID)
	s.Equal(req.Format, expectedFilmLog.Format)
	s.Equal(*req.Negative, *expectedFilmLog.Negative)
	s.Equal(*req.Brand, *expectedFilmLog.Brand)
	s.Equal(*req.ISO, *expectedFilmLog.ISO)
	s.Equal(*req.PurchaseDate, *expectedFilmLog.PurchaseDate)
	s.Equal(*req.LoadDate, *expectedFilmLog.LoadDate)
	s.Equal(req.Notes, expectedFilmLog.Notes)

	actualFilmLog, err := s.filmService.GetFilmLog(context.TODO(), 1)
	s.NoError(err)

	s.Equal(req.UserID, actualFilmLog.UserID)
	s.Equal(req.Format, actualFilmLog.Format)
	s.Equal(*req.Negative, *actualFilmLog.Negative)
	s.Equal(*req.Brand, *actualFilmLog.Brand)
	s.Equal(*req.ISO, *actualFilmLog.ISO)
	s.Equal(*req.PurchaseDate, *actualFilmLog.PurchaseDate)
	s.Equal(*req.LoadDate, *actualFilmLog.LoadDate)
	s.Equal(req.Notes, actualFilmLog.Notes)
}

func (s *filmServiceTestSuite) TestUpdateFilmWithInvalidReq() {
	s.Error(s.filmService.UpdateFilmLog(context.TODO(), film.UpdateFilmLogRequest{}),
		"empty req")
	s.Error(s.filmService.UpdateFilmLog(context.TODO(), film.UpdateFilmLogRequest{
		FilmLogID: 9999,
	}), "without user id")
	s.Error(s.filmService.UpdateFilmLog(context.TODO(), film.UpdateFilmLogRequest{
		FilmLogID: 9999,
		UserID:    util.Pointer(uint(9999)),
	}), "user not found")
}

func (s *filmServiceTestSuite) TestUpdateFilm() {
	s.userRepo.CreateUser(context.TODO(), &domain.User{
		ID: 2,
	})

	s.filmRepo.CreateFilmLog(context.TODO(), &domain.FilmLog{
		ID:           2,
		UserID:       2,
		Format:       domain.FilmFormat135,
		Negative:     util.Pointer(true),
		Brand:        util.Pointer("Kodak"),
		ISO:          util.Pointer(400),
		PurchaseDate: util.Pointer(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		LoadDate:     util.Pointer(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		Notes:        "",
	})

	req := film.UpdateFilmLogRequest{
		FilmLogID: 2,
		UserID:    util.Pointer(uint(2)),

		Format: util.Pointer(domain.FilmFormat120),
		Notes:  util.Pointer("updated"),
	}

	s.NoError(s.filmService.UpdateFilmLog(context.TODO(), req))

	actualFilmLog, err := s.filmService.GetFilmLog(context.TODO(), 2)
	s.NoError(err)

	// unchanged
	s.Equal(*req.Format, actualFilmLog.Format)
	s.Equal(true, *actualFilmLog.Negative)
	s.Equal("Kodak", *actualFilmLog.Brand)
	s.Equal(400, *actualFilmLog.ISO)
	s.Equal("2020-01-01", actualFilmLog.PurchaseDate.Format("2006-01-02"))
	s.Equal("2020-01-01", actualFilmLog.LoadDate.Format("2006-01-02"))

	// changed
	s.Equal(*req.UserID, actualFilmLog.UserID)
	s.Equal(*req.Notes, actualFilmLog.Notes)
}

func TestFilmServiceSuite(t *testing.T) {
	suite.Run(t, new(filmServiceTestSuite))
}
