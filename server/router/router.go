package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hugo/go-socketio/user"
)

var r *gin.Engine

func InitRouter(userCtr *user.Controller) {
	r = gin.Default()

	r.POST("/signup", userCtr.CreateUser)
	r.POST("/login", userCtr.Login)
	r.GET("/logout", userCtr.Logout)
}

func Start(addr string) error {
	return r.Run(addr)
}
