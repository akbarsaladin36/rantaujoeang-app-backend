package services

import (
	"rantaujoeang-app-backend/helpers"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/models"
	"rantaujoeang-app-backend/repositories"
	"time"
)

type MessageService interface {
	FindMessagesService() ([]models.Message, error)
	FindMessagesByUserIdService(user_uuid string) ([]models.Message, error)
	FindMessageService(message_code string) (models.Message, error)
	CreateMessageService(createMessageInput inputs.CreateMessageInput, currentUser map[string]string) (models.Message, error)
	ReplyMessageService(message_code string, replyMessageInput inputs.ReplyMessageInput, currentUser map[string]string) (models.Message, error)
	UpdateMessageService(message_code string, updateMessageInput inputs.UpdateMessageInput, currentUser map[string]string) (models.Message, error)
	DeleteMessageService(message_code string) (models.Message, error)
}

type messageService struct {
	messageRepository repositories.MessageRepository
}

func NewMessageService(messageRepository repositories.MessageRepository) *messageService {
	return &messageService{messageRepository}
}

func (ms *messageService) FindMessagesService() ([]models.Message, error) {
	messages, err := ms.messageRepository.FindMessages()

	return messages, err
}

func (ms *messageService) FindMessagesByUserIdService(user_uuid string) ([]models.Message, error) {
	messages, err := ms.messageRepository.FindMessagesByUserId(user_uuid)

	return messages, err
}

func (ms *messageService) FindMessageService(message_code string) (models.Message, error) {
	message, err := ms.messageRepository.FindMessage(message_code)

	return message, err
}

func (ms *messageService) CreateMessageService(createMessageInput inputs.CreateMessageInput, currentUser map[string]string) (models.Message, error) {
	messageCode := helpers.GenerateSlug(createMessageInput.MessageTitle)

	message := models.Message{
		SenderId:                   currentUser["user_uuid"],
		ReceiverId:                 createMessageInput.ReceiverId,
		MessageCode:                messageCode,
		MessageTitle:               createMessageInput.MessageTitle,
		MessageDescription:         createMessageInput.MessageDescription,
		MessageStatusCd:            "sent",
		MessageCreatedDate:         time.Now(),
		MessageCreatedUserUuid:     currentUser["user_uuid"],
		MessageCreatedUserUsername: currentUser["user_username"],
	}

	newMessage, err := ms.messageRepository.CreateMessage(message)

	return newMessage, err
}

func (ms *messageService) ReplyMessageService(message_code string, replyMessageInput inputs.ReplyMessageInput, currentUser map[string]string) (models.Message, error) {
	checkMessage, _ := ms.messageRepository.FindMessage(message_code)

	message := models.Message{
		SenderId:                   currentUser["user_uuid"],
		ReceiverId:                 checkMessage.ReceiverId,
		MessageCode:                checkMessage.MessageCode,
		MessageTitle:               "Reply For: " + replyMessageInput.MessageTitle,
		MessageDescription:         replyMessageInput.MessageDescription,
		MessageStatusCd:            "reply",
		MessageCreatedDate:         time.Now(),
		MessageCreatedUserUuid:     currentUser["user_uuid"],
		MessageCreatedUserUsername: currentUser["user_username"],
	}

	replyMessage, err := ms.messageRepository.CreateMessage(message)

	return replyMessage, err
}

func (ms *messageService) UpdateMessageService(message_code string, updateMessageInput inputs.UpdateMessageInput, currentUser map[string]string) (models.Message, error) {
	checkMessage, _ := ms.messageRepository.FindMessage(message_code)

	checkMessage.MessageCode = helpers.GenerateSlug(updateMessageInput.MessageTitle)
	checkMessage.MessageTitle = updateMessageInput.MessageTitle
	checkMessage.MessageDescription = updateMessageInput.MessageDescription
	checkMessage.MessageUpdatedDate = time.Now()
	checkMessage.MessageUpdatedUserUsername = currentUser["user_uuid"]
	checkMessage.MessageUpdatedUserUsername = currentUser["user_username"]

	updateMessage, err := ms.messageRepository.UpdateMessage(checkMessage)

	return updateMessage, err
}

func (ms *messageService) DeleteMessageService(message_code string) (models.Message, error) {
	checkMessage, _ := ms.messageRepository.FindMessage(message_code)

	deleteMessage, err := ms.messageRepository.DeleteMessage(checkMessage)

	return deleteMessage, err
}
