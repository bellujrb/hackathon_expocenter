package storages

import (
	"net/http"
	"viva/internal/interfaces"
	storages "viva/internal/storages/service"

	"github.com/gin-gonic/gin"
)

// @Summary Create cashback
// @Description Create new cashback
// @Tags Storage
// @Accept json
// @Produce json
// @Param request body interfaces.StorageCashBack true "Data for modify the cashback "ID = Storage in /storages" "
// @Success 200 {object} string "Cashback Create"
// @Router /api/modify-cashback [put]
func ModifyCashback(c *gin.Context) {
	var storage interfaces.StorageCashBack

	if err := c.ShouldBindJSON(&storage); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}

	if err := storages.PutCashBack(c, storage); err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Storage update successfully")
	c.Status(http.StatusOK)

}

// @Summary Create Pushed
// @Description Create new Pushed
// @Tags Storage
// @Accept json
// @Produce json
// @Param request body interfaces.StorageBuying true "Data for modify the pushed "ID = Storage in /storages" "
// @Success 200 {object} string "Pushed Create"
// @Router /api/modify-pushed [put]
func ModifyPushed(c *gin.Context) {
	var storage interfaces.StorageBuying

	if err := c.ShouldBindJSON(&storage); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}

	if err := storages.PutBuying(c, storage); err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Storage update successfully")
	c.Status(http.StatusOK)

}

// @Summary CashBacksUse
// @Description Create new CashBacksUse
// @Tags Storage
// @Accept json
// @Produce json
// @Param request body interfaces.StorageUse true "Data for modify the Used "ID = Storage in /storages" "
// @Success 200 {object} string "Used Create"
// @Router /api/modify-use [put]
func ModifyUse(c *gin.Context) {
	var storage interfaces.StorageUse

	if err := c.ShouldBindJSON(&storage); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusNotAcceptable)
		return
	}

	if err := storages.PutCashBackUses(c, storage); err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Storage update successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All Storage
// @Description Retrieve all storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} []db.Store "List of All Storage"
// @Router /api/storages [get]
func PullAllStorage(c *gin.Context) {
	storages.GetAllStorages(c)
}

// @Summary Retrieve Storage by ID
// @Description Retrieve an storage by its ID
// @Tags Storage
// @Accept json
// @Produce json
// @Param id path int true "Storage ID"
// @Success 200 {object} db.Store "Storage Details"
// @Router /api/storage/{id} [get]
func PullStorageId(c *gin.Context) {
	idParam := c.Param("id")
	storages.PullStorageIdService(c, idParam)

}
