package db

import "time"

type OpeningHours struct {
	Monday    string `json:"monday" xorm:"monday"`
	Tuesday   string `json:"tuesday" xorm:"tuesday"`
	Wednesday string `json:"wednesday" xorm:"wednesday"`
	Thursday  string `json:"thursday" xorm:"thursday"`
	Friday    string `json:"friday" xorm:"friday"`
	Saturday  string `json:"saturday" xorm:"saturday"`
	Sunday    string `json:"sunday" xorm:"sunday"`
}

type Store struct {
	ID           int          `json:"id" xorm:"pk autoincr 'id' INT"`
	Storage      string       `json:"storage" xorm:"storage"`
	Name         string       `json:"name" xorm:"name"`
	Address      string       `json:"address" xorm:"address"`
	City         string       `json:"city" xorm:"city"`
	State        string       `json:"state" xorm:"state"`
	ZipCode      string       `json:"zipCode" xorm:"zip_code"`
	Phone        string       `json:"phone" xorm:"phone"`
	Email        string       `json:"email" xorm:"email"`
	Website      string       `json:"website" xorm:"website"`
	Category     string       `json:"category" xorm:"category"`
	OpeningHours OpeningHours `json:"openingHours" xorm:"extends"`
	CashBack     float64      `json:"cashback" xorm:"cashback"`
	Buying       int64        `json:"buying" xorm:"buying"`
	CashBacksUse int64        `json:"cashBacksUse" xorm:"cashBacks"`
}

type Transaction struct {
	ID          int64     `json:"id" xorm:"pk autoincr 'id' INT"`
	CustomerID  int64     `xorm:"customer_id"`
	StoreID     int64     `xorm:"store_id"`
	Date        time.Time `xorm:"date"`
	Amount      float64   `xorm:"amount"`
	Type        string    `xorm:"type"`
	Description string    `xorm:"description"`
}

type CustomerTransactions struct {
	ID           int           `json:"id" xorm:"pk autoincr 'id' INT"`
	CustomerID   string        `xorm:"customer_id"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `xorm:"-"`
}

type Customer struct {
	ID               int64     `json:"id" xorm:"pk autoincr 'id'"`
	FirstName        string    `json:"firstName" xorm:"first_name"`
	LastName         string    `json:"lastName" xorm:"last_name"`
	Address          string    `json:"address" xorm:"address"`
	City             string    `json:"city" xorm:"city"`
	State            string    `json:"state" xorm:"state"`
	ZipCode          string    `json:"zipCode" xorm:"zip_code"`
	Phone            string    `json:"phone" xorm:"phone"`
	Email            string    `json:"email" xorm:"email"`
	DateOfBirth      time.Time `json:"dateOfBirth" xorm:"date_of_birth"`
	RegistrationDate time.Time `json:"registrationDate" xorm:"registration_date"`
	Xp               float64   `json:"xp" xorm:"xp"`
}

type CustomerPreference struct {
	ID         int    `json:"id" xorm:"pk autoincr 'id'"`
	CustomerID int    `json:"customerId" xorm:"customer_id"`
	Preference string `json:"preference" xorm:"preference"`
}

type Marketplace struct {
	ID          int    `json:"id" xorm:"pk autoincr 'id'"`
	Product     string `json:"product" xorm:"product"`
	Value       int64  `json:"value" xorm:"value"`
	Description string `json:"description" xorm:"description"`
	Img         string `json:"img" xorm:"img"`
}
