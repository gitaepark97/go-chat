package main

import (
	"log"

	"github.com/hugo/go-socketio/db"
	"github.com/hugo/go-socketio/router"
	"github.com/hugo/go-socketio/user"
	"github.com/hugo/go-socketio/ws"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userCtr := user.NewController(userSvc)

	hub := ws.NewHub()
	wsCtr := ws.NewController(hub)
	go hub.Run()

	router.InitRouter(userCtr, wsCtr)
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("could not start router: %s", err)
	}
}
