package user

import (
	"net/http"
	"strconv"
	"viva/internal/interfaces"
	user "viva/internal/user/service"

	"github.com/gin-gonic/gin"
)

// @Summary Create New Constumer
// @Description Create New Constumer
// @Tags Mkt
// @Accept json
// @Produce json
// @Param request body interfaces.CustomerInput true "Data for create new Constumer"
// @Success 200 {object} db.Customer "Return for creat Constumer"
// @Router /api/constumer [post]
func CreateConstumer(c *gin.Context) {
	var item interfaces.CustomerInput
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}
	if err := user.CreateUserService(c, item); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "New item created successfully")
	c.Status(http.StatusOK)
}

// @Summary Retrieve Constumer by ID
// @Description Retrieve an Constumer by its ID
// @Tags Mkt
// @Accept json
// @Produce json
// @Param id path int true "Constumer ID"
// @Success 200 {object} db.Customer "Constumer Details"
// @Router /api/constumer/{id} [get]
func GetOneConstumer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	user.PullCustomerId(c, id)
}

// @Summary Retrieve Constumer
// @Description Retrieve Constumer
// @Tags Constumer
// @Accept json
// @Produce json
// @Success 200 {object} db.Customer "List of Constumer"
// @Router /api/all-constumer [get]
func AllConstumer(c *gin.Context) {
	user.PullAllConstumer(c)
}

// @Summary Sync Constumer for external api
// @Description Retrieve Constumer
// @Tags Constumer
// @Accept json
// @Produce json
// @Success 200 {object} db.Customer "List of Constumer"
// @Router /api/sync-customers [get]
func SyncConstumers(c *gin.Context) {
	user.SyncExternalCustomers(c)
}
