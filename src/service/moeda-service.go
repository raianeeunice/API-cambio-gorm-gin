package service

import (
	"cambioo/src/dto"
	"cambioo/src/entity"
	"cambioo/src/repository"
	"log"

	"github.com/mashingan/smapping"
)

type MoedaService interface {
	InsertMoeda(moeda dto.MoedaDTO) (entity.Moeda, error)
	AllMoedas() []entity.Moeda
}

type moedaService struct {
	moedaRepository repository.MoedaRepository
}

//NewMoedatoService cria uma nova instância de moedaService
func NewMoedaService(moedaRepo repository.MoedaRepository) MoedaService {
	return &moedaService{
		moedaRepository: moedaRepo,
	}
}

func (service *moedaService) InsertMoeda(moeda dto.MoedaDTO) (entity.Moeda, error) {
	moedaCreate := entity.Moeda{}
	err := smapping.FillStruct(&moedaCreate, smapping.MapFields(&moeda))
	if err != nil {
		log.Fatalf("Failed maps %v: ", err)
	}
	res, err := service.moedaRepository.InsertMoeda(moedaCreate)
	return res, err
}

func (service *moedaService) AllMoedas() []entity.Moeda {
	return service.moedaRepository.FindAllMoedas()
}
