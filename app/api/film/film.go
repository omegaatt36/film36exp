package film

import (
	"strconv"
	"time"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/service/film"
	"github.com/omegaatt36/film36exp/util"

	"github.com/gin-gonic/gin"
)

// Controller is a film controller
type Controller struct {
	filmService *film.Service
}

// NewController create a new film controller
func NewController(filmService *film.Service) *Controller {
	return &Controller{filmService: filmService}
}

type filmLog struct {
	ID           uint    `json:"id"`
	UserID       uint    `json:"user_id"`
	Format       string  `json:"format"`
	Negative     *bool   `json:"negative"`
	Brand        *string `json:"brand"`
	ISO          *int    `json:"iso"`
	PurchaseDate *int64  `json:"purchase_date"`
	LoadDate     *int64  `json:"load_date"`
	Notes        string  `json:"notes"`
}

func (log *filmLog) fromDomain(domainLog *domain.FilmLog) {
	log.ID = domainLog.ID
	log.UserID = domainLog.UserID
	log.Format = domainLog.Format.String()
	log.Negative = domainLog.Negative
	log.Brand = domainLog.Brand
	log.ISO = domainLog.ISO
	if domainLog.PurchaseDate != nil {
		log.PurchaseDate = util.Pointer(domainLog.PurchaseDate.Unix())
	}
	if domainLog.LoadDate != nil {
		log.LoadDate = util.Pointer(domainLog.LoadDate.Unix())
	}
	log.Notes = domainLog.Notes
}

type createFilmLogRequest struct {
	UserID       uint              `json:"user_id" binding:"required"`
	Format       domain.FilmFormat `json:"format" binding:"required"`
	Negative     *bool             `json:"negative"`
	Brand        *string           `json:"brand"`
	ISO          *int              `json:"iso"`
	PurchaseDate *time.Time        `json:"purchase_date"`
	LoadDate     *time.Time        `json:"load_date"`
	Notes        string            `json:"notes"`
}

func (x *Controller) CreateFilmLog(c *gin.Context) {
	var req createFilmLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.filmService.CreateFilmLog(c.Request.Context(), film.CreateFilmLogRequest{
		UserID:       req.UserID,
		Format:       req.Format,
		Negative:     req.Negative,
		Brand:        req.Brand,
		ISO:          req.ISO,
		PurchaseDate: req.PurchaseDate,
		LoadDate:     req.LoadDate,
		Notes:        req.Notes,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// GetFilmLog get a film log
func (x *Controller) GetFilmLog(c *gin.Context) {
	filmLogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	domainFilmLog, err := x.filmService.GetFilmLog(c.Request.Context(), uint(filmLogID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp filmLog
	resp.fromDomain(domainFilmLog)

	c.JSON(200, resp)
}

type updateFilmLogRequest struct {
	UserID       *uint              `json:"user_id"`
	Format       *domain.FilmFormat `json:"format"`
	Negative     *bool              `json:"negative"`
	Brand        *string            `json:"brand"`
	ISO          *int               `json:"iso"`
	PurchaseDate *time.Time         `json:"purchase_date"`
	LoadDate     *time.Time         `json:"load_date"`
	Notes        *string            `json:"notes"`
}

// UpdateFilmLog update a film log
func (x *Controller) UpdateFilmLog(c *gin.Context) {
	filmLogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req updateFilmLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.filmService.UpdateFilmLog(c.Request.Context(), film.UpdateFilmLogRequest{
		FilmLogID:    uint(filmLogID),
		UserID:       req.UserID,
		Format:       req.Format,
		Negative:     req.Negative,
		Brand:        req.Brand,
		ISO:          req.ISO,
		PurchaseDate: req.PurchaseDate,
		LoadDate:     req.LoadDate,
		Notes:        req.Notes,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// DeleteFilmLog delete a film log
func (x *Controller) DeleteFilmLog(c *gin.Context) {
	filmLogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.filmService.DeleteFilmLog(c.Request.Context(), uint(filmLogID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}
