package main

import (
	"cambioo/src/config"
	"cambioo/src/controller"
	"cambioo/src/repository"
	"cambioo/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	depositoRepository repository.DepositoRepository = repository.NewDepositoRepository(db)
	depositoService    service.DepositoService       = service.NewdepositoService(depositoRepository)
	depositoController controller.DepositoController = controller.NewdepositoController(depositoService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	depositoRoutes := r.Group("api/depositos")
	{
		depositoRoutes.GET("/", depositoController.AllDepositos)
		depositoRoutes.POST("/", depositoController.InsertDeposito)
		depositoRoutes.GET("/:id", depositoController.FindDepositoByID)
	}

	saldoRoutes := r.Group("api/saldo")
	{
		saldoRoutes.GET("/", depositoController.FindSaldoTotal)
		saldoRoutes.GET("/:moeda", depositoController.Cambio)
	}

	r.Run()
} 