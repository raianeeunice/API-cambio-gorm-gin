package service

import (
	"cambioo/src/dto"
	"cambioo/src/entity"
	"cambioo/src/repository"
	"errors"
	"log"
	"math"

	"github.com/mashingan/smapping"
)

//depositoService é um contrato do que depositoService pode fazer
type DepositoService interface {
	InsertDeposito(b dto.DepositoCreateDTO) entity.Depositos
	AllDepositos() []entity.Depositos
	FindDepositoByID(depositoID uint64) entity.Depositos
	FindSaldoTotal() float64
	ConverterMoeda(moeda string) float64
	ValidarMoedaService(moeda string) error
	converter(valor float64, moeda string) float64
	calculaCambio(moedaCotacao float64) float64
	ajusteCasaDecimal(valor float64) float64
}

type depositoService struct {
	depositoRepository repository.DepositoRepository
	moedaRepository repository.MoedaRepository
}

//NewdepositoService cria uma nova instância de depositoService
func NewdepositoService(depositoRepo repository.DepositoRepository, moedaRepo repository.MoedaRepository) DepositoService {
	return &depositoService{
		depositoRepository: depositoRepo,
		moedaRepository: moedaRepo,
	}
}

func (service *depositoService) InsertDeposito(valorDeposito dto.DepositoCreateDTO) entity.Depositos {
	deposito := entity.Depositos{}
	err := smapping.FillStruct(&deposito, smapping.MapFields(&valorDeposito))
	if err != nil {
		log.Fatalf("Failed maps %v: ", err)
	}
	res := service.depositoRepository.InsertDeposito(deposito)
	return res
}

func (service *depositoService) AllDepositos() []entity.Depositos {
	return service.depositoRepository.FindAllDepositos()
}

func (service *depositoService) FindDepositoByID(depositoID uint64) entity.Depositos {
	return service.depositoRepository.FindDepositoByID(depositoID)
}

func (service *depositoService) FindSaldoTotal() float64{
	var saldo = service.depositoRepository.FindSaldoTotal()
	return service.ajusteCasaDecimal(saldo)
}

func (service *depositoService) ConverterMoeda(moeda string) float64 {
	saldoTotal := service.depositoRepository.FindSaldoTotal()
	conversao := service.converter(saldoTotal, moeda)
	return service.ajusteCasaDecimal(conversao)
}

func (service *depositoService) ValidarMoedaService(moeda string) error {
	validar := service.moedaRepository.FindBySigla(moeda)
	if validar == ""{
		return errors.New("sigla inexistnte")
	}
	return nil
}

func (service *depositoService) converter(valor float64, moeda string) float64 {
	cotacao := service.moedaRepository.FindCotacaoBySigla(moeda);
	return valor/service.calculaCambio(cotacao)
}

func (service *depositoService) calculaCambio(cotacao float64) float64 {
	iof := 0.0638
	taxaCambio := 0.16
	valorFinal := cotacao + (cotacao * iof) + (cotacao * taxaCambio)
	return service.ajusteCasaDecimal(valorFinal)
}

func (service *depositoService) ajusteCasaDecimal(valor float64) float64{
	return math.Floor(valor*100)/100
}