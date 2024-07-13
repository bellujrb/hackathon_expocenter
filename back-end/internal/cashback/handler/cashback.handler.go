package cashback

import (
	"fmt"
	"net/http"
	cashback "viva/internal/cashback/service"
	"viva/internal/interfaces"

	"github.com/gin-gonic/gin"
)

// @Summary Calculation Cashback Total
// @Description Calculation Cashback Total
// @Tags Cashback
// @Accept json
// @Produce json
// @Param request body interfaces.Purchase true "For add new cashbacks in wallet to consumer"
// @Success 200 {object} string "Used Create"
// @Router /api/calc-cashback [post]
func CalculationCashBack(c *gin.Context) {
	var pushed interfaces.Purchase

	if err := c.ShouldBindJSON(&pushed); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}
	cash, err := cashback.CalculationCashBackService(c, pushed)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusNotAcceptable)
	}
	responseMessage := fmt.Sprintf("Your new cashback: %.2f", cash)
	c.Set("Response", responseMessage)
	c.Status(http.StatusOK)

}

// @Summary Remove Cashback Total
// @Description Remove Cashback Total
// @Tags Cashback
// @Accept json
// @Produce json
// @Param request body interfaces.Purchase true "For add new cashbacks in wallet to consumer"
// @Success 200 {object} string "Used Create"
// @Router /api/remove-cashback [post]
func RemoveCashBack(c *gin.Context) {
	var pushed interfaces.Purchase

	if err := c.ShouldBindJSON(&pushed); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}
	cash, err := cashback.CalculationCashBackService(c, pushed)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusNotAcceptable)
	}
	responseMessage := fmt.Sprintf("Your new cashback: %.2f", cash)
	c.Set("Response", responseMessage)
	c.Status(http.StatusOK)

}
