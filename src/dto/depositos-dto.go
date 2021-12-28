package dto

import "errors"

type DepositoCreateDTO struct {
	ValorDeposito float64 `json:"valor_deposito"`
}

// validar é a função responsável por validar o campo ValorDepositado
func (deposito *DepositoCreateDTO) Validar() error {
	if deposito.ValorDeposito <= 0 {
		return errors.New("o valor depositado precisa se maior do que zero e não pode estar em branco")
	}
	return nil
}
