package models

import "time"

type Session struct {
	SessionId           int    `json:"session_id" gorm:"primaryKey"`
	SessionToken        string `json:"session_token" gorm:"type:text"`
	SessionUserUUID     string `json:"session_user_uuid" gorm:"type:varchar(200)"`
	SessionUserUsername string `json:"session_user_username" gorm:"type:varchar(100)"`
	SessionUserRole     string `json:"session_user_role" gorm:"type:varchar(30)"`
	SessionStartAt      time.Time
	SessionExpiredAt    time.Time
	SessionCreatedAt    time.Time
	SessionUpdateAt     time.Time
}
