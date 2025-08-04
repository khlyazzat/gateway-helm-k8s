package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Profile struct {
	bun.BaseModel `bun:"table:profiles,alias:p"`
	ID            int64 `bun:",pk,autoincrement"`
	UserID        int64
	FirstName     string
	LastName      string
	Age           int
	Address       string
	Phone         string
	CreatedAt     time.Time
	UpdatedAt     bun.NullTime
	DeletedAt     bun.NullTime `bun:",soft_delete,nullzero"`
}
