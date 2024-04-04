package film

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/service/film"
	"github.com/omegaatt36/film36exp/util"
)

type photo struct {
	ID           uint     `json:"id"`
	FilmLogID    uint     `json:"film_log_id"`
	Aperture     *float64 `json:"aperture"`
	ShutterSpeed *string  `json:"shutter_speed"`
	Date         *int64   `json:"date"`
	Description  *string  `json:"description"`
	Tags         []string `json:"tags"`
	Location     *string  `json:"location"`
}

func (x *photo) fromDomain(domainPhoto *domain.Photo) {
	x.ID = domainPhoto.ID
	x.FilmLogID = domainPhoto.FilmLogID
	x.Aperture = domainPhoto.Aperture
	x.ShutterSpeed = domainPhoto.ShutterSpeed
	if domainPhoto.Date != nil {
		x.Date = util.Pointer(domainPhoto.Date.Unix())
	}
	x.Description = domainPhoto.Description
	x.Tags = domainPhoto.Tags
	x.Location = domainPhoto.Location
}

type createPhotoRequest struct {
	FilmLogID    uint `validate:"required"`
	Aperture     *float64
	ShutterSpeed *string
	Date         *int64
	Description  *string
	Tags         []string
	Location     *string
}

func (x *Controller) CreatePhoto(c *gin.Context) {
	var req createPhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var date *time.Time
	if req.Date != nil {
		date = util.Pointer(time.Unix(*req.Date, 0))
	}

	if err := x.filmService.CreatePhoto(c.Request.Context(), film.CreatePhotoRequest{
		FilmLogID:    req.FilmLogID,
		Aperture:     req.Aperture,
		ShutterSpeed: req.ShutterSpeed,
		Date:         date,
		Description:  req.Description,
		Tags:         req.Tags,
		Location:     req.Location,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// GetPhoto get a photo
func (x *Controller) GetPhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	domainPhoto, err := x.filmService.GetPhoto(c.Request.Context(), uint(photoID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp photo
	resp.fromDomain(domainPhoto)

	c.JSON(200, resp)
}

type updatePhotoRequest struct {
	FilmLogID    *uint      `json:"film_log_id"`
	Aperture     *float64   `json:"aperture"`
	ShutterSpeed *string    `json:"shutter_speed"`
	Date         *time.Time `json:"date"`
	Description  *string    `json:"description"`
	Tags         []string   `json:"tags"`
	Location     *string    `json:"location"`
}

// UpdatePhoto update a photo
func (x *Controller) UpdatePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req updatePhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.filmService.UpdatePhoto(c.Request.Context(), film.UpdatePhotoRequest{
		PhotoID:      uint(photoID),
		FilmLogID:    req.FilmLogID,
		Aperture:     req.Aperture,
		ShutterSpeed: req.ShutterSpeed,
		Date:         req.Date,
		Description:  req.Description,
		Tags:         req.Tags,
		Location:     req.Location,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// DeletePhoto delete a photo
func (x *Controller) DeletePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.filmService.DeletePhoto(c.Request.Context(), uint(photoID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}
