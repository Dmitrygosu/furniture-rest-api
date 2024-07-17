package model

type Furniture struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Fabricator string `json:"fabricator"`
	Height     uint32 `json:"height"`
	Width      uint32 `json:"width"`
	Length     uint32 `json:"length"`
}

// HasEmptyFields- проверка нан пустые поля
func (f *Furniture) HasEmptyFields() bool {
	return f.Name == "" || f.Fabricator == "" || f.Height == 0 || f.Width == 0 || f.Length == 0
}
