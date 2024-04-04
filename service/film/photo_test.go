package film_test

import (
	"context"
	"time"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/service/film"
	"github.com/omegaatt36/film36exp/util"
)

func (s *filmServiceTestSuite) TestCreatePhotoWithInvalidReq() {
	s.Error(s.filmService.CreatePhoto(context.TODO(), film.CreatePhotoRequest{}),
		"empty req")
	s.Error(s.filmService.CreatePhoto(context.TODO(), film.CreatePhotoRequest{
		FilmLogID: 9999,
	}), "film log not found")
}

func (s *filmServiceTestSuite) TestCreatePhoto() {
	s.NoError(s.filmRepo.CreateFilmLog(context.TODO(), &domain.FilmLog{
		ID:     3,
		UserID: 1,
	}))

	req := film.CreatePhotoRequest{
		FilmLogID:    3,
		Aperture:     util.Pointer(2.8),
		ShutterSpeed: util.Pointer("1/500"),
		Date:         util.Pointer(time.Date(2022, 3, 14, 12, 0, 0, 0, time.UTC)),
		Description:  util.Pointer("Test description"),
		Tags:         []string{"test", "photo"},
		Location:     util.Pointer("Somewhere"),
	}
	s.NoError(s.filmService.CreatePhoto(context.TODO(), req))

	expectedPhoto, err := s.filmRepo.GetPhoto(context.TODO(), 1)
	s.NoError(err)

	s.Equal(req.FilmLogID, expectedPhoto.FilmLogID)
	s.Equal(req.Aperture, expectedPhoto.Aperture)
	s.Equal(req.ShutterSpeed, expectedPhoto.ShutterSpeed)
	s.Equal(req.Date, expectedPhoto.Date)
	s.Equal(req.Description, expectedPhoto.Description)
	s.Equal(req.Tags, expectedPhoto.Tags)
	s.Equal(req.Location, expectedPhoto.Location)

	actualPhoto, err := s.filmService.GetPhoto(context.TODO(), 1)
	s.NoError(err)

	s.Equal(req.FilmLogID, actualPhoto.FilmLogID)
	s.Equal(req.Aperture, actualPhoto.Aperture)
	s.Equal(req.ShutterSpeed, actualPhoto.ShutterSpeed)
	s.Equal(req.Date, actualPhoto.Date)
	s.Equal(req.Description, actualPhoto.Description)
	s.Equal(req.Tags, actualPhoto.Tags)
	s.Equal(req.Location, actualPhoto.Location)
}

func (s *filmServiceTestSuite) TestUpdatePhotoWithInvalidReq() {
	s.Error(s.filmService.UpdatePhoto(context.TODO(), film.UpdatePhotoRequest{}),
		"empty req")
	s.Error(s.filmService.UpdatePhoto(context.TODO(), film.UpdatePhotoRequest{
		PhotoID: 9999,
	}), "photo not found")
}

func (s *filmServiceTestSuite) TestUpdatePhoto() {
	s.NoError(s.filmRepo.CreatePhoto(context.TODO(), &domain.Photo{
		ID:           4,
		FilmLogID:    3,
		Aperture:     util.Pointer(2.8),
		ShutterSpeed: util.Pointer("1/500"),
		Date:         util.Pointer(time.Date(2022, 3, 14, 12, 0, 0, 0, time.UTC)),
		Description:  util.Pointer("Test description"),
		Tags:         []string{"test", "photo"},
		Location:     util.Pointer("Somewhere"),
	}))

	req := film.UpdatePhotoRequest{
		PhotoID: 4,
		// FilmLogID:    3,
		Aperture:     util.Pointer(3.5),
		ShutterSpeed: util.Pointer("1/100"),
		Date:         util.Pointer(time.Date(2022, 3, 15, 12, 0, 0, 0, time.UTC)),
		Description:  util.Pointer("Updated description"),
		Tags:         []string{"updated", "photo"},
		Location:     util.Pointer("Updated location"),
	}
	s.NoError(s.filmService.UpdatePhoto(context.TODO(), req))

	actualPhoto, err := s.filmService.GetPhoto(context.TODO(), 4)
	s.NoError(err)

	s.Equal(req.PhotoID, actualPhoto.ID)
	s.Equal(req.Aperture, actualPhoto.Aperture)
	s.Equal(req.ShutterSpeed, actualPhoto.ShutterSpeed)
	s.Equal(req.Date, actualPhoto.Date)
	s.Equal(req.Description, actualPhoto.Description)
	s.Equal(req.Tags, actualPhoto.Tags)
}
