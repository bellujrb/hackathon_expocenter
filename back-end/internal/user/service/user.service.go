package user

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"viva/db"
	"viva/external"
	"viva/internal/interfaces"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func CreateUserService(c *gin.Context, data interfaces.CustomerInput) error {
	user := db.Customer{
		FirstName:        data.FirstName,
		LastName:         data.LastName,
		Address:          data.Address,
		City:             data.City,
		State:            data.State,
		ZipCode:          data.ZipCode,
		Phone:            data.Phone,
		Email:            data.Email,
		DateOfBirth:      data.DateOfBirth,
		RegistrationDate: time.Now(),
		Xp:               0,
	}
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	if err := db.Create(engine.(*xorm.Engine), &user); err != nil {
		return err
	}
	return nil
}

func PullAllConstumer(c *gin.Context) {
	var constumer []db.Customer
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &constumer); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", constumer)
	c.Status(http.StatusOK)
}

func PullCustomerId(c *gin.Context, id int) {
	var constumer db.Customer
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &constumer, int64(id))
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
	c.Set("Response", constumer)
	c.Status(http.StatusOK)
}

func SyncExternalCustomers(c *gin.Context) {
	var customers []db.Customer
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Response", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	// Fetch external customer details
	res, err := external.FetchUsersDetails()
	if err != nil {
		c.Set("Response", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// Get all customers from the database
	if err := db.GetAll(engine.(*xorm.Engine), &customers); err != nil {
		c.Set("Response", "Error retrieving customers")
		c.Status(http.StatusInternalServerError)
		return
	}

	apiCount := len(res)
	dbCount := len(customers)

	if apiCount > dbCount {
		for _, extCustomer := range res {
			extCustomerID, err := strconv.ParseInt(extCustomer.ID, 10, 64)
			if err != nil {
				c.Set("Response", err)
				c.Status(http.StatusInternalServerError)
				return
			}

			var dbCustomer db.Customer
			found, err := db.GetByID(engine.(*xorm.Engine), &dbCustomer, extCustomerID)
			if err != nil {
				c.Set("Response", err)
				c.Status(http.StatusInternalServerError)
				return
			}

			if !found {
				dbCustomer = db.Customer{
					FirstName:        extCustomer.FirstName,
					LastName:         extCustomer.LastName,
					Address:          extCustomer.Address,
					City:             extCustomer.City,
					State:            extCustomer.State,
					ZipCode:          extCustomer.ZipCode,
					Phone:            extCustomer.Phone,
					Email:            extCustomer.Email,
					DateOfBirth:      extCustomer.DateOfBirth,
					RegistrationDate: extCustomer.RegistrationDate,
					Xp:               0,
				}

				err = db.Create(engine.(*xorm.Engine), &dbCustomer)
				if err != nil {
					c.Set("Response", err)
					c.Status(http.StatusInternalServerError)
					return
				}
			}
		}
	}

	c.Set("Response", customers)
	c.Status(http.StatusOK)
}
