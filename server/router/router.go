package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hugo/go-socketio/user"
	"github.com/hugo/go-socketio/ws"
)

var r *gin.Engine

func InitRouter(userCtr *user.Controller, wsCtr *ws.Controller) {
	r = gin.Default()

	r.POST("/signup", userCtr.CreateUser)
	r.POST("/login", userCtr.Login)
	r.GET("/logout", userCtr.Logout)

	r.POST("/ws/createRoom", wsCtr.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsCtr.JoinRoom)
	r.GET("/ws/getRooms", wsCtr.GetRooms)
	r.GET("/ws/getClients/:roomId", wsCtr.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
