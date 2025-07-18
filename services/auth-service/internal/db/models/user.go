package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64 `bun:",pk,autoincrement"`
	Name          string
	Password      string
	Email         string
	Age           int
	CreatedAt     time.Time
	UpdatedAt     bun.NullTime
	DeletedAt     bun.NullTime `bun:",soft_delete,nullzero"`
}
