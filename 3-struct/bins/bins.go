package bins

type BinList []Bin

// Структура для хранения информации о бине
type Bin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
