package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tickets/internal/application/interfaces"
	"tickets/internal/domain/entities"
)

type UserHandler struct {
	userUseCase interfaces.UserUseCase
}

func NewUserHandler(userUseCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.userUseCase.Create(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")
	ctx := c.Request.Context()
	user, err := u.userUseCase.FindByUuid(ctx, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.userUseCase.Update(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	ctx := c.Request.Context()
	err := u.userUseCase.Delete(ctx, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
