package entity

import (
	"time"
)

type Pack struct {
	ID        uint32
	CreatedAt time.Time
	Size      uint32
}

func NewPack(size int) Pack {
	return Pack{
		CreatedAt: time.Now().UTC(),
		Size:      uint32(size),
	}
}
