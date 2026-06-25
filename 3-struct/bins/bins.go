package bins

import (
	"time"

	"github.com/google/uuid"
)

type BinList []Bin
type Bins interface {
	NewBin(name string, private bool) *Bin
}

func NewBins() Bins {
	return &binsImpl{}
}

type binsImpl struct {
	binList BinList
}

// Структура для хранения информации о бине
type Bin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Private   bool      `json:"private"`
}

// Функция для создания нового бина
func (bins *binsImpl) NewBin(name string, private bool) *Bin {
	return &Bin{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}
}
