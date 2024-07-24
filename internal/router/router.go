package router

import (
	"copilot/pkg/copilot/use-model/controllers"
	"copilot/pkg/copilot/use-model/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	abastecimentoService := services.NewAbastecimentoService(client, "cibus-copilot", "fills")
	abastecimentoController := controllers.NewAbastecimentoController(abastecimentoService)
	r.GET("/ping", abastecimentoController.GetAbastecimentos)

	return r
}
