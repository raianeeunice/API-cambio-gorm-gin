package entity

type Moeda struct {
	Sigla     string  `json:"sigla" gorm:"primaryKey"`
	Descricao string  `json:"descricao" gorm:"unique"`
	Cotacao   float64 `json:"cotacao"`
}
