package main

import (
	"fmt"
	"log"

	"github.com/moroz/oauth-starter/models"
)

func main() {
	db := models.ConnectToDB()
	db.MustExec("delete from users")
	user, err := models.CreateUser(db, models.CreateUserParams{Email: "user@example.com"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", user)
}
