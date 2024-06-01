package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-CURD-GoLang/service/user/model"
	"strconv"
)

type Business interface {
	CreateUser(ctx context.Context, user *model.CreateUserReq) error
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*model.User, error)
}

type Api struct {
	business Business
}

func NewApi(business Business) *Api {
	return &Api{business: business}
}

func (a *Api) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.CreateUserReq

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := a.business.CreateUser(c.Request.Context(), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func (a *Api) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := a.business.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "User deleted successfully",
		})
	}
}
func (a *Api) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		user, err := a.business.GetUser(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}
