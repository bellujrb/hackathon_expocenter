package mkt

import (
	"errors"
	"net/http"
	"viva/db"
	"viva/internal/interfaces"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func CreateItemService(c *gin.Context, data interfaces.MktInput) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}
	inputMkt := db.Marketplace{
		Product:     data.Product,
		Value:       data.Value,
		Description: data.Description,
		Img:         data.Img,
	}
	if err := db.Create(engine.(*xorm.Engine), &inputMkt); err != nil {
		return err
	}
	return nil
}

func PullAllItem(c *gin.Context) {
	var mkt []db.Marketplace
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &mkt); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", mkt)
	c.Status(http.StatusOK)
}

func PullControlId(c *gin.Context, id int) {
	var mkt db.Marketplace
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &mkt, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving item")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "item not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", mkt)
	c.Status(http.StatusOK)
}
