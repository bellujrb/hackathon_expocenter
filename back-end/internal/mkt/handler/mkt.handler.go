package mkt

import (
	"net/http"
	"strconv"
	"viva/internal/interfaces"
	mkt "viva/internal/mkt/service"

	"github.com/gin-gonic/gin"
)

// @Summary Create New Item For marketPlace
// @Description Create New Item For marketPlace
// @Tags Mkt
// @Accept json
// @Produce json
// @Param request body interfaces.MktInput true "Data for create new Item"
// @Success 200 {object} db.Marketplace "Return for creat item"
// @Router /api/mkt [post]
func CreateItem(c *gin.Context) {
	var item interfaces.MktInput
	if err := c.ShouldBindJSON(&item); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}
	if err := mkt.CreateItemService(c, item); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "New item created successfully")
	c.Status(http.StatusOK)
}

// @Summary Retrieve Item by ID
// @Description Retrieve an Item by its ID
// @Tags Mkt
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} db.Marketplace "Item Details"
// @Router /api/mkt/{id} [get]
func GetOneItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	mkt.PullControlId(c, id)
}

// @Summary Retrieve Item Control
// @Description Retrieve Item Event
// @Tags Mkt
// @Accept json
// @Produce json
// @Success 200 {object} db.Marketplace "List of Item Event"
// @Router /api/all-item [get]
func AllItem(c *gin.Context) {
	mkt.PullAllItem(c)
}
