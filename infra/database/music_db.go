package database

import (
	"database/sql"

	"github.com/vitorconti/loader-pagination/internal/entity"
)

type Music struct {
	DB *sql.DB
}

func NewMusic(db *sql.DB) *Music {
	return &Music{DB: db}
}

func (p *Music) FindAll(page, limit int) ([]entity.Music, error) {
	var products []entity.Music
	var err error
	if page != 0 && limit != 0 {
		//todo
	}
	return products, err
}
