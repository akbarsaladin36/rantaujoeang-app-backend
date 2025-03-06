package models

import "time"

type Dorm struct {
	DormId                  int       `json:"dorm_id" gorm:"primaryKey"`
	DormCode                string    `json:"dorm_code" gorm:"type:varchar(150)"`
	DormName                string    `json:"dorm_name" gorm:"type:varchar(150)"`
	DormPriceAmount         string    `json:"dorm_price_amount" gorm:"type:varchar(200)"`
	DormQuantityAmount      string    `json:"dorm_quantity_amount" gorm:"type:varchar(150)"`
	DormAddress             string    `json:"dorm_address" gorm:"type:text"`
	DormDescription         string    `json:"dorm_description" gorm:"type:text"`
	DormPhoneNumber         string    `json:"dorm_phone_number" gorm:"type:varchar(30)"`
	DormStatusCd            string    `json:"dorm_status_cd" gorm:"type:varchar(30)"`
	DormCreatedDate         time.Time `json:"dorm_created_date"`
	DormCreatedUserUuid     string    `json:"dorm_created_user_uuid" gorm:"type:varchar(200)"`
	DormCreatedUserUsername string    `json:"dorm_created_user_username" gorm:"type:varchar(100)"`
	DormUpdatedDate         time.Time `json:"dorm_updated_date"`
	DormUpdatedUserUuid     string    `json:"dorm_updated_user_uuid" gorm:"type:varchar(200)"`
	DormUpdatedUserUsername string    `json:"dorm_updated_user_username" gorm:"type:varchar(100)"`
}
