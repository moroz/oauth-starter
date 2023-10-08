package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/uuidv7-go"
)

type User struct {
	Id         uuidv7.UUID
	Email      string
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type CreateUserParams struct {
	Email string
}

func CreateUser(db *sqlx.DB, params CreateUserParams) (*User, error) {
	uuid := uuidv7.Generate()
	var user User
	err := db.QueryRowx("insert into users (id, email) values ($1, $2) returning id, email, inserted_at, updated_at", uuid.Dump(), params.Email).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(db *sqlx.DB, id uuidv7.UUID) (*User, error) {
	var user User
	err := db.QueryRowx("select id, email, inserted_at, updated_at from users where id = $1", id).StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserById(db *sqlx.DB, id uuidv7.UUID) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}
