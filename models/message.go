package models

import "time"

type Message struct {
	MessageId                  int    `json:"message_id" gorm:"primaryKey"`
	SenderId                   string `json:"sender_id" gorm:"type:varchar(200)"`
	ReceiverId                 string `json:"receiver_id" gorm:"type:varchar(200)"`
	MessageCode                string `json:"message_code" gorm:"type:varchar(150)"`
	MessageTitle               string `json:"message_title" gorm:"type:varchar(150)"`
	MessageDescription         string `json:"message_description" gorm:"type:text"`
	MessageStatusCd            string `json:"message_status_cd" gorm:"type:varchar(30)"`
	MessageCreatedDate         time.Time
	MessageCreatedUserUuid     string `json:"message_created_user_uuid" gorm:"type:varchar(200)"`
	MessageCreatedUserUsername string `json:"message_created_user_username" gorm:"type:varchar(100)"`
	MessageUpdatedDate         time.Time
	MessageUpdatedUserUuid     string `json:"message_updated_user_uuid" gorm:"type:varchar(200)"`
	MessageUpdatedUserUsername string `json:"message_updated_user_username" gorm:"type:varchar(100)"`
}
