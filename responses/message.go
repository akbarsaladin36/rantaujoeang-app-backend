package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type MessageResponse struct {
	SenderId           string `json:"sender_id"`
	ReceiverId         string `json:"receiver_id"`
	MessageCode        string `json:"message_code"`
	MessageTitle       string `json:"message_title"`
	MessageDescription string `json:"message_description"`
	MessageStatusCd    string `json:"message_status_cd"`
}

type CreateMessageResponse struct {
	SenderId                   string    `json:"sender_id"`
	ReceiverId                 string    `json:"receiver_id"`
	MessageCode                string    `json:"message_code"`
	MessageTitle               string    `json:"message_title"`
	MessageDescription         string    `json:"message_description"`
	MessageStatusCd            string    `json:"message_status_cd"`
	MessageCreatedDate         time.Time `json:"message_created_date"`
	MessageCreatedUserUuid     string    `json:"message_created_user_uuid"`
	MessageCreatedUserUsername string    `json:"message_created_user_username"`
}

type UpdateMessageResponse struct {
	SenderId                   string    `json:"sender_id"`
	ReceiverId                 string    `json:"receiver_id"`
	MessageCode                string    `json:"message_code"`
	MessageTitle               string    `json:"message_title"`
	MessageDescription         string    `json:"message_description"`
	MessageStatusCd            string    `json:"message_status_cd"`
	MessageUpdatedDate         time.Time `json:"message_updated_date"`
	MessageUpdatedUserUuid     string    `json:"message_updated_user_uuid"`
	MessageUpdatedUserUsername string    `json:"message_updated_user_username"`
}

func GetMessageResponse(messageRsps models.Message) MessageResponse {
	return MessageResponse{
		SenderId:           messageRsps.SenderId,
		ReceiverId:         messageRsps.ReceiverId,
		MessageCode:        messageRsps.MessageCode,
		MessageTitle:       messageRsps.MessageTitle,
		MessageDescription: messageRsps.MessageDescription,
		MessageStatusCd:    messageRsps.MessageStatusCd,
	}
}

func GetCreateMessageResponse(messageRsps models.Message) CreateMessageResponse {
	return CreateMessageResponse{
		SenderId:                   messageRsps.SenderId,
		ReceiverId:                 messageRsps.ReceiverId,
		MessageCode:                messageRsps.MessageCode,
		MessageTitle:               messageRsps.MessageTitle,
		MessageDescription:         messageRsps.MessageDescription,
		MessageStatusCd:            messageRsps.MessageStatusCd,
		MessageCreatedDate:         messageRsps.MessageCreatedDate,
		MessageCreatedUserUuid:     messageRsps.MessageCreatedUserUuid,
		MessageCreatedUserUsername: messageRsps.MessageCreatedUserUsername,
	}
}

func GetUpdateMessageResponse(messageRsps models.Message) UpdateMessageResponse {
	return UpdateMessageResponse{
		SenderId:                   messageRsps.SenderId,
		ReceiverId:                 messageRsps.ReceiverId,
		MessageCode:                messageRsps.MessageCode,
		MessageTitle:               messageRsps.MessageTitle,
		MessageDescription:         messageRsps.MessageDescription,
		MessageStatusCd:            messageRsps.MessageStatusCd,
		MessageUpdatedDate:         messageRsps.MessageUpdatedDate,
		MessageUpdatedUserUuid:     messageRsps.MessageUpdatedUserUuid,
		MessageUpdatedUserUsername: messageRsps.MessageUpdatedUserUsername,
	}
}
