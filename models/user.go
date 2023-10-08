package models

import (
	"time"

	"github.com/moroz/oauth-starter/pkg/uuidv7"
)

type User struct {
	Id         uuidv7.UUID
	Email      string
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
