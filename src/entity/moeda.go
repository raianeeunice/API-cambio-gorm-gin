package entity

type Moeda struct {
	Sigla     string  `json:"sigla" gorm:"primaryKey"`
	Descricao string  `json:"descricao" gorm:"not null"`
	Cotacao   float64 `json:"cotacao" gorm:"not null"`
}
