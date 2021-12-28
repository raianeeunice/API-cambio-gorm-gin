package repository

import (
	"cambioo/src/entity"

	"gorm.io/gorm"
)

//DepositoRepository é um contrato com o que o depositoRepository pode fazer com o banco de dados
type DepositoRepository interface {
	InsertDeposito(deposito entity.Depositos) entity.Depositos
	FindAllDepositos() []entity.Depositos
	FindDepositoByID(depositoID uint64) entity.Depositos
	FindSaldoTotal() float64
}

type depositoConnection struct {
	connection *gorm.DB
}

//NewDepositoRepository cria uma instancia de DepositoRepository
func NewDepositoRepository(dbConn *gorm.DB) DepositoRepository {
	return &depositoConnection{
		connection: dbConn,
	}
}

func (db *depositoConnection) InsertDeposito(deposito entity.Depositos) entity.Depositos {
	db.connection.Save(&deposito)
	db.connection.Find(&deposito)
	return deposito
}

func (db *depositoConnection) FindAllDepositos() []entity.Depositos {
	var deposito []entity.Depositos
	db.connection.Find(&deposito)
	return deposito
}

func (db *depositoConnection) FindDepositoByID(depositoID uint64) entity.Depositos {
	var deposito entity.Depositos
	db.connection.Find(&deposito, depositoID)
	return deposito
}

func (db *depositoConnection) FindSaldoTotal() float64 {
	var saldo float64
	db.connection.Table("depositos").Select("sum(valor_deposito) as saldo").Scan(&saldo) //query nativa
	
	return saldo
}