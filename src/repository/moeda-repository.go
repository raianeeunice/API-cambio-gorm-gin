package repository

import (
	"cambioo/src/entity"
	"gorm.io/gorm"
)

type MoedaRepository interface {
	InsertMoeda(moeda entity.Moeda) (entity.Moeda, error)
	FindAllMoedas() []entity.Moeda
	FindBySigla(moedaSigla string) string
	FindCotacaoBySigla(moedaSigla string) float64
}

type moedaConnection struct {
	connection *gorm.DB
}

//NewDepositoRepository cria uma instancia de MoedaRepository
func NewMoedaRepository(dbConn *gorm.DB) MoedaRepository {
	return &moedaConnection{
		connection: dbConn,
	}
}

func (db *moedaConnection) InsertMoeda(moeda entity.Moeda) (entity.Moeda, error) {
	err := db.connection.Save(&moeda).Error

	return moeda, err
}

func (db *moedaConnection) FindAllMoedas() []entity.Moeda {
	var moeda []entity.Moeda
	db.connection.Find(&moeda)
	return moeda
}

func (db *moedaConnection) FindBySigla(moedaSigla string) string {
	var moeda string
	db.connection.Table("moedas").Select("sigla").Where("sigla = ?", moedaSigla).Scan(&moeda)
	return moeda
}

func (db *moedaConnection) FindCotacaoBySigla(moedaSigla string) float64 {
	var cotacao float64
	db.connection.Table("moedas").Select("cotacao").Where("sigla = ?", moedaSigla).Scan(&cotacao)
	return cotacao
}
