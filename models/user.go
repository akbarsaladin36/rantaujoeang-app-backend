package models

import "time"

type User struct {
	UserId                       int    `json:"user_id" gorm:"primaryKey"`
	UserUUID                     string `json:"user_uuid" gorm:"type:varchar(200)"`
	UserUsername                 string `json:"user_username" gorm:"type:varchar(100)"`
	UserPassword                 string `json:"user_password" gorm:"type:varchar(180)"`
	UserEmail                    string `json:"user_email" gorm:"type:varchar(150)"`
	UserFirstName                string `json:"user_first_name" gorm:"type:varchar(150)"`
	UserLastName                 string `json:"user_last_name" gorm:"type:varchar(150)"`
	UserAddress                  string `json:"user_address" gorm:"type:text"`
	UserRole                     string `json:"user_role" gorm:"type:varchar(30)"`
	UserStatusCd                 string `json:"user_status_cd" gorm:"type:varchar(30)"`
	UserBalanceTransactionAmount string `json:"user_balance_transaction_amount" gorm:"type:varchar(255)"`
	UserCreatedDate              time.Time
	UserCreatedUserUuid          string `json:"user_created_user_uuid" gorm:"type:varchar(200)"`
	UserCreatedUsername          string `json:"user_created_user_username" gorm:"type:varchar(100)"`
	UserUpdatedDate              time.Time
	UserUpdatedUserUuid          string `json:"user_updated_user_uuid" gorm:"type:varchar(200)"`
	UserUpdatedUsername          string `json:"user_updated_user_username" gorm:"type:varchar(100)"`
}
