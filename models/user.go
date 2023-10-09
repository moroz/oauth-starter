package models

import (
	"database/sql"
	"fmt"
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

const USER_FIELDS = "id, email, inserted_at, updated_at"
const insertUserQuery = "insert into users (id, email) values ($1, $2) returning " + USER_FIELDS

func CreateUser(db *sqlx.DB, params CreateUserParams) (*User, error) {
	uuid := uuidv7.Generate()
	var user User
	err := db.QueryRowx(insertUserQuery, uuid.Dump(), params.Email).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: %w", err)
	}
	return &user, nil
}

const getUserByEmailQuery = "select " + USER_FIELDS + " from users where email = $1"

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var user User
	err := db.Get(&user, getUserByEmailQuery, email)
	switch err {
	case nil:
		return &user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

const getUserByIdQuery = "select " + USER_FIELDS + " from users where id = $1"

func GetUserById(db *sqlx.DB, id string) (*User, error) {
	var user User
	err := db.Get(&user, getUserByIdQuery, id)
	switch err {
	case nil:
		return &user, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func DeleteUserById(db *sqlx.DB, id uuidv7.UUID) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}

func GetOrCreateUserByEmail(db *sqlx.DB, email string) (*User, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return CreateUser(db, CreateUserParams{Email: email})
	}

	return user, nil
}
