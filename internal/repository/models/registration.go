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
	ID                 int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationId     int64  `json:"registration_id"`
	ThaiName           string `json:"thai_name"`
	EngName            string `json:"eng_name"`
	Category           string `json:"category"`
	SubCategory        string `json:"sub_category"`
	SkuCount           string `json:"sku_count"`
	Exclusivity        bool   `json:"exclusivity"`
	MarkoSale          bool   `json:"marko_sale"`
	LotussSale         bool   `json:"lotuss_sale"`
	TargetCustomerType string `json:"target_customer_type"`
	AvailableChannels  string `json:"available_channels"`
	USP                string `json:"usp"`
	AvailableMarket    string `json:"available_market"`
}

type RegistrationItemImage struct {
	ID                 string `json:"id" gorm:"primaryKey,autoIncrement"`
	RegistrationItemId int64  `json:"registration_item_id"`
	FileName           string `gorm:"not null"`
	MimeType           string `gorm:"not null"`
	Data               []byte `gorm:"type:bytea;not null"`
}
