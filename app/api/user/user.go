package user

import (
	"strconv"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/service/user"

	"github.com/gin-gonic/gin"
)

// Controller is a user controller
type Controller struct {
	userService *user.Service
}

// NewController create a new user controller
func NewController(userService *user.Service) *Controller {
	return &Controller{userService: userService}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser create a new user
func (x *Controller) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.userService.CreateUser(c.Request.Context(), user.CreateUserRequest(req)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

type userDetail struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *userDetail) fromDomain(domainUser *domain.User) {
	u.ID = domainUser.ID
	u.Name = domainUser.Name
	u.Email = domainUser.Email
}

// GetUser get a user
func (x *Controller) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	domainUser, err := x.userService.GetUser(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp userDetail
	resp.fromDomain(domainUser)

	c.JSON(200, resp)
}

type updateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// UpdateUser update a user
func (x *Controller) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.userService.UpdateUser(c.Request.Context(), user.UpdateUserRequest{
		UserID:   uint(userID),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

// DeleteUser delete a user
func (x *Controller) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := x.userService.DeleteUser(c.Request.Context(), uint(userID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}
