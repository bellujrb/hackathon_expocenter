package interfaces

import "time"

type StorageCashBack struct {
	ID       string  `json:"id"`
	Cashback float64 `json:"cashback"`
}
type StorageBuying struct {
	ID     string `json:"id"`
	Buying int64  `json:"buying"`
}
type StorageUse struct {
	ID           string `json:"id"`
	CashBacksUse int64  `json:"buying"`
}

type Purchase struct {
	Value   float64 `json:"value"`
	Email   string  `json:"email"`
	Storage string  `json:"storage"`
}

type MktInput struct {
	Product     string `json:"product"`
	Value       int64  `json:"value"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type CustomerInput struct {
	FirstName   string    `json:"firstName" xorm:"first_name"`
	LastName    string    `json:"lastName" xorm:"last_name"`
	Address     string    `json:"address" xorm:"address"`
	City        string    `json:"city" xorm:"city"`
	State       string    `json:"state" xorm:"state"`
	ZipCode     string    `json:"zipCode" xorm:"zip_code"`
	Phone       string    `json:"phone" xorm:"phone"`
	Email       string    `json:"email" xorm:"email"`
	DateOfBirth time.Time `json:"dateOfBirth" xorm:"date_of_birth"`
}
