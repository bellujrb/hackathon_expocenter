package cashback

import (
	"errors"
	"net/http"
	"viva/db"
	"viva/internal/interfaces"
	storages "viva/internal/storages/service"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func CalculationCashBackService(c *gin.Context, input interfaces.Purchase) (float64, error) {
	var storage db.Store
	var user db.Customer
	var transactions db.CustomerTransactions
	engine, exists := c.Get("db")
	if !exists {
		return 0, errors.New("database connection not found")
	}
	inputPurchase := interfaces.StorageBuying{
		ID:     input.Storage,
		Buying: 1,
	}
	if err := storages.PutBuying(c, inputPurchase); err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	found, err := db.GetByStorageId(engine.(*xorm.Engine), &storage, input.Storage)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, err
	}
	found, err = db.GetByEmail(engine.(*xorm.Engine), &user, input.Email)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	if !found {
		return 0, errors.New("customer not found")
	}
	found, err = db.GetByID(engine.(*xorm.Engine), &transactions, user.ID)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	if !found {
		return 0, errors.New("customer not found")
	}
	cashback := input.Value * storage.CashBack

	updateValues := map[string]interface{}{
		"Balance": cashback + transactions.Balance,
	}
	updateXp := map[string]interface{}{
		"Xp": user.Xp + (5 * cashback),
	}
	if err := db.Update(engine.(*xorm.Engine), transactions, updateValues); err != nil {
		return 0, err
	}
	if err := db.Update(engine.(*xorm.Engine), user, updateXp); err != nil {
		return 0, err
	}
	return cashback, nil
}

func RemoveCashBack(c *gin.Context, input interfaces.Purchase) (float64, error) {
	var storage db.Store
	var user db.Customer
	var transactions db.CustomerTransactions
	engine, exists := c.Get("db")
	if !exists {
		return 0, errors.New("database connection not found")
	}
	inputPurchase := interfaces.StorageUse{
		ID:           input.Storage,
		CashBacksUse: 1,
	}
	if err := storages.PutCashBackUses(c, inputPurchase); err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	found, err := db.GetByStorageId(engine.(*xorm.Engine), &storage, input.Storage)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, err
	}
	found, err = db.GetByEmail(engine.(*xorm.Engine), &user, input.Email)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	if !found {
		return 0, errors.New("customer not found")
	}
	found, err = db.GetByID(engine.(*xorm.Engine), &transactions, user.ID)
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return 0, err
	}
	if !found {
		return 0, errors.New("customer not found")
	}
	new := input.Value - transactions.Balance
	if err := db.Update(engine.(*xorm.Engine), transactions, new); err != nil {
		return 0, err
	}
	return new, nil
}
