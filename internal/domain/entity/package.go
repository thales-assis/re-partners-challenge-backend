package entity

import (
	"time"
)

type Package struct {
	ID        uint32
	CreatedAt time.Time
	Size      uint32
}

func NewPackage(size int) Package {
	return Package{
		CreatedAt: time.Now().UTC(),
		Size:      uint32(size),
	}
}
