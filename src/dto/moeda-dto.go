package dto

import (
	"errors"
	"strings"
)

type MoedaDTO struct {
	Sigla     string  `json:"sigla"`
	Descricao string  `json:"descricao"`
	Cotacao   float64 `json:"cotacao"`
}

func (moeda *MoedaDTO) ValidarMoeda() error {
	if erro := moeda.verificarValorDigitado(); erro != nil {
		return erro
	}

	if erro := moeda.formatarMoeda(); erro != nil {
		return erro
	}

	return nil
}

func (moeda *MoedaDTO) verificarValorDigitado() error {
	if moeda.Sigla == "" {
		return errors.New("a sigla é obrigatória e não pode estar em branco")
	}
	if moeda.Descricao == "" {
		return errors.New("a descrição é obrigatória e não pode estar em branco")
	}
	if moeda.Cotacao <= 0 {
		return errors.New("a cotação é obrigatória e não pode estar em branco")
	}

	return nil
}

func (moeda *MoedaDTO) formatarMoeda() error {
	moeda.Sigla= strings.TrimSpace(moeda.Sigla)
	moeda.Descricao = strings.TrimSpace(moeda.Descricao)

	return nil
}