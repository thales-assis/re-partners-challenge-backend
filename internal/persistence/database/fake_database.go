package database

import (
	"time"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type FakeDatabase struct {
	CountIDs uint32
	Records  map[uint32]entity.Pack
}

type FakeDatabaseTotalItems struct {
	CountIDs uint32                 `json:"count_ids"`
	Records  []FakeDatabasePackItem `json:"records"`
}

type FakeDatabasePackItem struct {
	ID        uint32    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Size      uint32    `json:"size"`
}
