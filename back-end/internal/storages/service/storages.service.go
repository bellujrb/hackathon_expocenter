package storages

import (
	"errors"
	"net/http"
	"viva/db"
	"viva/external"
	"viva/internal/interfaces"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func PutCashBack(c *gin.Context, input interfaces.StorageCashBack) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}
	updateValues := map[string]interface{}{
		"cashback": input.Cashback,
	}
	if err := db.UpdateStorage(engine.(*xorm.Engine), input.ID, updateValues); err != nil {
		return err
	}
	return nil
}

func PutCashBackUses(c *gin.Context, input interfaces.StorageUse) error {
	var storage db.Store
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}
	found, err := db.GetByStorageId(engine.(*xorm.Engine), &storage, input.ID)
	if err != nil {
		return err
	}
	if !found {
		return err
	}

	updateValues := map[string]interface{}{
		"buying": input.CashBacksUse + storage.CashBacksUse,
	}

	if err := db.UpdateStorage(engine.(*xorm.Engine), input.ID, updateValues); err != nil {
		return err
	}
	return nil
}

func PutBuying(c *gin.Context, input interfaces.StorageBuying) error {
	var storage db.Store
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}
	found, err := db.GetByStorageId(engine.(*xorm.Engine), &storage, input.ID)
	if err != nil {
		return err
	}
	if !found {
		return err
	}

	updateValues := map[string]interface{}{
		"buying": input.Buying + storage.Buying,
	}

	if err := db.UpdateStorage(engine.(*xorm.Engine), input.ID, updateValues); err != nil {
		return err
	}
	return nil
}

func GetAllStorages(c *gin.Context) {
	var storage []db.Store
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Response", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	res, err := external.FetchStoreDetails()
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &storage); err != nil {
		c.Set("Response", "Error retrieving storage")
		c.Status(http.StatusInternalServerError)
		return
	}
	apiCount := len(res)
	dbCount := len(storage)

	if apiCount > dbCount {
		for _, store := range res {
			var storeMigration db.Store
			found, err := db.GetByStorageId(engine.(*xorm.Engine), &storeMigration, store.ID)
			if err != nil {
				c.Set("Response", err)
				c.Status(http.StatusInternalServerError)
				return
			}

			if !found {
				err = db.Create(engine.(*xorm.Engine), storeMigration)
				if err != nil {
					c.Set("Response", err)
					c.Status(http.StatusInternalServerError)
					return
				}
			}
		}
	}

	c.Set("Response", storage)
	c.Status(http.StatusOK)
}

func PullStorageIdService(c *gin.Context, id string) {
	var storage interfaces.StorageCashBack
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByStorageId(engine.(*xorm.Engine), &storage, id)
	if err != nil {
		c.Set("Error", "Error retrieving Storage")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "Storage not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", storage)
	c.Status(http.StatusOK)
}
