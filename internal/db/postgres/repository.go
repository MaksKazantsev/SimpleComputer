package postgres

import (
	"college/internal/db"
	"gorm.io/gorm"
)

var _ db.Repository = &Postgres{}

type Postgres struct {
	*gorm.DB
}

func NewRepository(conn *gorm.DB) *Postgres {
	return &Postgres{conn}
}
