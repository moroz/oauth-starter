package models

import (
	"time"

	"github.com/GoWebProd/uuid7"
)

type User struct {
	Id         uuid7.UUID
	Email      string
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
