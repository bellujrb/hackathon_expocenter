package external

import "time"

type OpeningHours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

type Store struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Address      string       `json:"address"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	ZipCode      string       `json:"zipCode"`
	Phone        string       `json:"phone"`
	Email        string       `json:"email"`
	Website      string       `json:"website"`
	Category     string       `json:"category"`
	OpeningHours OpeningHours `json:"openingHours"`
}

type Customer struct {
	ID               string    `json:"id" xorm:"pk 'id'"`
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
	Preferences      []string  `json:"preferences" xorm:"-"`
}
