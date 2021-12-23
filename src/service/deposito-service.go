package service

import (
	"cambioo/src/dto"
	"cambioo/src/entity"
	"cambioo/src/repository"
	"log"

	"github.com/mashingan/smapping"
)

//depositoService é um contrato do que depositoService pode fazer
type DepositoService interface {
	InsertDeposito(b dto.DepositoCreateDTO) entity.Depositos
	AllDepositos() []entity.Depositos
	FindDepositoByID(depositoID uint64) entity.Depositos
	FindSaldoTotal() float64
	Cambio() float64
}

type depositoService struct {
	depositoRepository repository.DepositoRepository
}

//NewdepositoService cria uma nova instância de depositoService
func NewdepositoService(depositoRepo repository.DepositoRepository) DepositoService {
	return &depositoService{
		depositoRepository: depositoRepo,
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
	return service.depositoRepository.FindSaldoTotal()
}

func (service *depositoService) Cambio() float64 {
	return service.depositoRepository.FindSaldoTotal()
}
