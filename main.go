package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/moroz/oauth-starter/config"
	"github.com/moroz/oauth-starter/controllers"
	"github.com/moroz/oauth-starter/models"
)

func main() {
	db := models.ConnectToDB()
	r := controllers.Routes(db)
	listener, err := net.Listen("tcp", config.LISTEN_ON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on %s\n", config.LISTEN_ON)
	http.Serve(listener, r)
}
