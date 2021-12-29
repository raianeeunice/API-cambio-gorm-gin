package main

import (
	"cambioo/src/config"
	"cambioo/src/controller"
	"cambioo/src/repository"
	"cambioo/src/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	depositoRepository repository.DepositoRepository = repository.NewDepositoRepository(db)
	depositoService    service.DepositoService       = service.NewdepositoService(depositoRepository, moedaRepository)
	depositoController controller.DepositoController = controller.NewdepositoController(depositoService)
	
	moedaRepository repository.MoedaRepository = repository.NewMoedaRepository(db)
	moedaService    service.MoedaService      = service.NewMoedaService(moedaRepository)
	moedaController controller.MoedaController = controller.NewMoedaController(moedaService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	 r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))

	depositoRoutes := r.Group("api/depositos")
	{
		depositoRoutes.GET("/", depositoController.AllDepositos)
		depositoRoutes.POST("/", depositoController.InsertDeposito)
		depositoRoutes.GET("/:id", depositoController.FindDepositoByID)
	}

	saldoRoutes := r.Group("api/saldo")
	{
		saldoRoutes.GET("/", depositoController.FindSaldoTotal)
		saldoRoutes.GET("/:moeda", depositoController.ConverterMoeda)
	}

	moedaRoutes := r.Group("api/moeda")
	{
		moedaRoutes.GET("/", moedaController.AllMoedas)
		moedaRoutes.POST("/", moedaController.InsertMoeda)
	}

	r.Run()
}
