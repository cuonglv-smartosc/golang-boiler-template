package models

import "time"

type Registration struct {
	ID                        int64     `json:"id" gorm:"primaryKey,autoIncrement"`
	ThaiCompanyName           string    `json:"thai_company_name"`
	EngCompanyName            string    `json:"eng_company_name"`
	ThaiCompanyAddress        string    `json:"thai_company_address"`
	EngCompanyAddress         string    `json:"eng_company_address"`
	Province                  string    `json:"province"`
	ZipCode                   string    `json:"zip_code"`
	CompanyEmail              string    `json:"company_email" gorm:"unique"`
	CompanyWebsite            string    `json:"company_website" gorm:"unique"`
	ContactNumber             string    `json:"contact_number"`
	TelephoneNumber           string    `json:"telephone_number" gorm:"unique"`
	AnnualRevenue             string    `json:"annual_revenue"`
	BusinessEstablishmentDate time.Time `json:"business_establishment_date"`
	TaxpayerNumber            string    `json:"taxpayer_number"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

type RegistrationDocument struct {
	ID             string `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationId int64  `json:"registration_item_id"`
	FileName       string `gorm:"not null"`
	MimeType       string `gorm:"not null"`
	Data           []byte `gorm:"type:bytea;not null"`
}

type RegistrationRepresentative struct {
	ID             int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationId int64  `json:"registration_id"`
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	ContactNumber  string `json:"contact_number"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RegistrationFactory struct {
	ID                    int64     `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationId        int64     `json:"registration_id"`
	Name                  string    `json:"name"`
	Address               string    `json:"address"`
	Province              string    `json:"province"`
	ZipCode               string    `json:"zip_code"`
	RegistrationNumber    string    `json:"registration_number"`
	LicenseExpirationDate time.Time `json:"license_expiration_date"`
	IsLicensed            bool      `json:"is_licensed"`
	Standards             string    `json:"standards"`
}

type RegistrationItem struct {
	ID                 int64
	RegistrationId     int64
	ThaiName           string
	EngName            string
	Category           string
	Subcategory        string
	SkuCount           string
	Exclusivity        bool
	MarkoSale          bool
	LotusSale          bool
	TargetCustomerType string
	AvailableChannels  string
	USP                string
	AvailableMarket    string
}

type RegistrationItemImage struct {
	ID                 string `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationItemId int64  `json:"registration_item_id"`
	FileName           string `gorm:"not null"`
	MimeType           string `gorm:"not null"`
	Data               []byte `gorm:"type:bytea;not null"`
}
