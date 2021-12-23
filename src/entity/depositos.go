package entity

import "time"

type Depositos struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	ValorDeposito float64        `json:"valorDeposito"`
	CreatedAt     time.Time      `json:"created"`
}