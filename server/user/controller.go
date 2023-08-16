package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service
}

func NewController(s Service) *Controller {
	return &Controller{
		Service: s,
	}
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var req CreateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.Service.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) Login(ctx *gin.Context) {
	var req LoginUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := c.Service.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	res := &LoginUserRes{
		Username: u.Username,
		ID:       u.ID,
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
