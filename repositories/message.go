package repositories

import (
	"rantaujoeang-app-backend/models"

	"gorm.io/gorm"
)

type MessageRepository interface {
	FindMessages() ([]models.Message, error)
	FindMessagesByUserId(user_uuid string) ([]models.Message, error)
	FindMessage(message_code string) (models.Message, error)
	CreateMessage(message models.Message) (models.Message, error)
	UpdateMessage(message models.Message) (models.Message, error)
	DeleteMessage(message models.Message) (models.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db}
}

func (mr *messageRepository) FindMessages() ([]models.Message, error) {
	var messages []models.Message

	err := mr.db.Find(&messages).Error

	return messages, err
}

func (mr *messageRepository) FindMessagesByUserId(user_uuid string) ([]models.Message, error) {
	var messages []models.Message

	err := mr.db.Where("sender_id = ?", user_uuid).Or("receiver_id = ?", user_uuid).Find(&messages).Error

	return messages, err
}

func (mr *messageRepository) FindMessage(message_code string) (models.Message, error) {
	var message models.Message

	err := mr.db.Where("message_code = ?", message_code).First(&message).Error

	return message, err
}

func (mr *messageRepository) CreateMessage(message models.Message) (models.Message, error) {
	err := mr.db.Create(&message).Error

	return message, err
}

func (mr *messageRepository) UpdateMessage(message models.Message) (models.Message, error) {
	err := mr.db.Save(&message).Error

	return message, err
}

func (mr *messageRepository) DeleteMessage(message models.Message) (models.Message, error) {
	err := mr.db.Delete(&message).Error

	return message, err
}
