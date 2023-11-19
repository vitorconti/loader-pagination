package database

import "github.com/vitorconti/loader-pagination/internal/entity"

type MusicInterface interface {
	FindAll(page, limit int) ([]entity.Music, error)
}
