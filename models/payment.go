package models

import "time"

type Payment struct {
	PaymentId                  int    `json:"payment_id" gorm:"primaryKey"`
	DormCode                   string `json:"dorm_code" gorm:"type:varchar(150)"`
	UserUuid                   string `json:"user_uuid" gorm:"type:varchar(200)"`
	PaymentCode                string `json:"payment_code" gorm:"type:varchar(150)"`
	PaymentAmount              string `json:"payment_amount" gorm:"type:varchar(200)"`
	PaymentQuantityAmount      string `json:"payment_quantity_amount" gorm:"type:varchar(150)"`
	PaymentDescription         string `json:"payment_description" gorm:"type:text"`
	PaymentTags                string `json:"payment_tags" gorm:"type:text"`
	PaymentStatusCd            string `json:"payment_status_cd" gorm:"type:varchar(30)"`
	PaymentCreatedDate         time.Time
	PaymentCreatedUserUuid     string `json:"payment_created_user_uuid" gorm:"type:varchar(200)"`
	PaymentCreatedUserUsername string `json:"payment_created_user_username" gorm:"type:varchar(100)"`
	PaymentUpdatedDate         time.Time
	PaymentUpdatedUserUuid     string `json:"payment_updated_user_uuid" gorm:"type:varchar(200)"`
	PaymentUpdatedUserUsername string `json:"payment_updated_user_username" gorm:"type:varchar(100)"`
}
