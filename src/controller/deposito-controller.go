package controller

import (
	"cambioo/src/dto"
	"cambioo/src/entity"
	"cambioo/src/helper"
	"cambioo/src/service"
	"cambioo/src/utils"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//depositoController é um contrato do que depositoController pode fazer
type DepositoController interface {
	InsertDeposito(context *gin.Context)
	AllDepositos(context *gin.Context)
	FindDepositoByID(context *gin.Context)
	FindSaldoTotal(context *gin.Context)
	Cambio(context *gin.Context)
}

type depositoController struct {
	depositoService service.DepositoService
}

//NewdepositoController cria uma nova instancia de BoookController
func NewdepositoController(depositoServ service.DepositoService) DepositoController {
	return &depositoController{
		depositoService: depositoServ,
	}
}


func (c *depositoController) InsertDeposito(context *gin.Context) {
	var depositoCreateDTO dto.DepositoCreateDTO
	errDTO := context.ShouldBind(&depositoCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Falha ao processar o pedido", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.depositoService.InsertDeposito(depositoCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *depositoController) AllDepositos(context *gin.Context) {
	var depositos []entity.Depositos = c.depositoService.AllDepositos()

	if (depositos == nil) {
		res := helper.BuildErrorResponse("Dados não encontrados", "Nenhum dado fornecido", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", depositos)
		context.JSON(http.StatusOK, res)
	}
}

func (c *depositoController) FindDepositoByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Nenhum id de parâmetro foi encontrado", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var deposito entity.Depositos = c.depositoService.FindDepositoByID(id)
	if (deposito == entity.Depositos{}) {
		res := helper.BuildErrorResponse("Dados não encontrados", "Nenhum dado com o id fornecido", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", deposito)
		context.JSON(http.StatusOK, res)
	}
}

func (c *depositoController) FindSaldoTotal(context *gin.Context){
	saldoTotal := c.depositoService.FindSaldoTotal()

	if (saldoTotal == 0) {
		res := helper.BuildErrorResponse("Dados não encontrados", "Nenhum saldo na conta", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", gin.H{
			"saldo_total": saldoTotal,
		})
		context.JSON(http.StatusOK, res)
	}
}

func (c *depositoController) Cambio(context *gin.Context){
	moeda := strings.ToUpper(context.Param("moeda"))
	saldoTotal := c.depositoService.FindSaldoTotal()
	conversao := utils.Converte(saldoTotal, moeda)
	conversao = math.Floor(conversao*100)/100

	if (conversao == 0) {
		res := helper.BuildErrorResponse("Dados não encontrados", "Nenhum dado com o tipo de moeda fornecido", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", gin.H{
			"saldo_total_convertido" : conversao,
		})
		context.JSON(http.StatusOK, res)
	}
}