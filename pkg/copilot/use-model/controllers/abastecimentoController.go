package controllers

import (
	"copilot/pkg/copilot/use-model/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AbastecimentoController struct {
	service *services.AbastecimentoService
}

func GetAbastecimento(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ponh",
	})
}

func NewAbastecimentoController(service *services.AbastecimentoService) *AbastecimentoController {
	return &AbastecimentoController{service}
}

func (ctrl *AbastecimentoController) GetAbastecimentos(c *gin.Context) {
	customer := c.Query("customer")

	abastecimentos, err := ctrl.service.GetAbastecimentoByCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, abastecimentos)
}
