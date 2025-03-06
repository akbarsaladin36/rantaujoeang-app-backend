package models

import "time"

type Invoice struct {
	InvoiceId                  int    `json:"invoice_id" gorm:"primaryKey"`
	PaymentCode                string `json:"payment_code" gorm:"type:varchar(150)"`
	DormCode                   string `json:"dorm_code" gorm:"type:varchar(150)"`
	UserUuid                   string `json:"user_uuid" gorm:"type:varchar(200)"`
	InvoiceCode                string `json:"invoice_code" gorm:"type:varchar(150)"`
	InvoiceAmount              string `json:"invoice_amount" gorm:"type:varchar(200)"`
	InvoiceQuantityAmount      string `json:"invoice_quantity_amount" gorm:"type:varchar(150)"`
	InvoiceDescription         string `json:"invoice_description" gorm:"type:text"`
	InvoiceTags                string `json:"invoice_tags" gorm:"type:text"`
	InvoiceStatusCd            string `json:"invoice_status_cd" gorm:"type:varchar(30)"`
	InvoiceCreatedDate         time.Time
	InvoiceCreatedUserUuid     string `json:"invoice_created_user_uuid" gorm:"type:varchar(200)"`
	InvoiceCreatedUserUsername string `json:"invoice_created_user_username" gorm:"type:varchar(100)"`
	InvoiceUpdatedDate         time.Time
	InvoiceUpdatedUserUuid     string `json:"invoice_updated_user_uuid" gorm:"type:varchar(200)"`
	InvoiceUpdatedUserUsername string `json:"invoice_updated_user_username" gorm:"type:varchar(100)"`
}
