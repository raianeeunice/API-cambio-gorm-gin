package controller

import (
	"cambioo/src/dto"
	"cambioo/src/entity"
	"cambioo/src/helper"
	"cambioo/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoedaController interface {
	InsertMoeda(context *gin.Context)
	AllMoedas(context *gin.Context)
}

type moedaController struct {
	moedaService service.MoedaService
}

//NewdepositoController cria uma nova instancia de BoookController
func NewMoedaController(moedaServ service.MoedaService) MoedaController {
	return &moedaController{
		moedaService: moedaServ,
	}
}

func (c *moedaController) InsertMoeda(context *gin.Context) {
	var moedaDTO dto.MoedaDTO
	errDTO := context.ShouldBind(&moedaDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Falha ao processar o pedido", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		err := moedaDTO.ValidarMoeda()
		if err != nil {
			res := helper.BuildErrorResponse("Valor inválido", err.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
		} else {
			result, err := c.moedaService.InsertMoeda(moedaDTO)
			if err != nil {
				res := helper.BuildErrorResponse("Moeda já cadastrada", err.Error(), helper.EmptyObj{})
				context.JSON(http.StatusBadRequest, res)
			} else {
				response := helper.BuildResponse(true, "OK", result)
				context.JSON(http.StatusCreated, response)
			}

		}
	}

}

func (c *moedaController) AllMoedas(context *gin.Context) {
	var moedas []entity.Moeda = c.moedaService.AllMoedas()

	if moedas == nil {
		res := helper.BuildErrorResponse("Dados não encontrados", "Nenhum dado fornecido", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", moedas)
		context.JSON(http.StatusOK, res)
	}
}
