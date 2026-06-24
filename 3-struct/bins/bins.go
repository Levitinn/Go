package bins

import (
	"time"

	"github.com/google/uuid"
)

type BinList []Bin

type Bin struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Private   bool
}

func NewBin(name string, private bool) *Bin {
	return &Bin{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}
}
