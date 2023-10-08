package main

import (
	"fmt"
	"log"

	"github.com/moroz/oauth-starter/models"
)

func main() {
	db := models.ConnectToDB()
	var user models.User
	err := db.QueryRowx("select id, email, inserted_at, updated_at from users order by id desc limit 1").StructScan(&user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", user)
}
