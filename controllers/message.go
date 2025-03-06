package controllers

import (
	"fmt"
	"net/http"
	"rantaujoeang-app-backend/inputs"
	"rantaujoeang-app-backend/middleware"
	"rantaujoeang-app-backend/responses"
	"rantaujoeang-app-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type messageController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) *messageController {
	return &messageController{messageService}
}

func (mc *messageController) FindMessagesController(c *gin.Context) {
	messages, errMessages := mc.messageService.FindMessagesService()

	if errMessages != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All messages data are empty! Please create a new message now",
		})
		return
	}

	var messagesRsps []responses.MessageResponse

	for _, message := range messages {
		messageRsps := responses.GetMessageResponse(message)

		messagesRsps = append(messagesRsps, messageRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All messages data are succesfully appeared!",
		"data":    messagesRsps,
	})
}

func (mc *messageController) FindMessagesByUserIdController(c *gin.Context) {
	currentUser := middleware.CurrentUser(c)

	messages, errMessages := mc.messageService.FindMessagesByUserIdService(currentUser["user_uuid"])

	if errMessages != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All messages data are empty! Please create a new message now",
		})
		return
	}

	var messagesRsps []responses.MessageResponse

	for _, message := range messages {
		messageRsps := responses.GetMessageResponse(message)

		messagesRsps = append(messagesRsps, messageRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All messages data are succesfully appeared!",
		"data":    messagesRsps,
	})
}

func (mc *messageController) FindMessageController(c *gin.Context) {
	message_code := c.Param("message_code")

	message, errMessage := mc.messageService.FindMessageService(message_code)

	if errMessage != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "A message data " + message_code + " are not found! Please try again!",
		})
		return
	}

	messageRsps := responses.GetMessageResponse(message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A message data " + message_code + " are succesfully appeared!",
		"data":    messageRsps,
	})
}

func (mc *messageController) CreateMessageController(c *gin.Context) {
	var createMessageInput inputs.CreateMessageInput

	errCreateMessageInput := c.ShouldBindJSON(&createMessageInput)

	if errCreateMessageInput != nil {
		errorMessages := []string{}
		for _, e := range errCreateMessageInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	createUser, err := mc.messageService.CreateMessageService(createMessageInput, currentUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when creating a message! Please try again!",
		})
		return
	}

	createUserRsps := responses.GetCreateMessageResponse(createUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new message is succesfully created!",
		"data":    createUserRsps,
	})
}

func (mc *messageController) ReplyMessageController(c *gin.Context) {
	message_code := c.Param("message_code")

	_, errMessage := mc.messageService.FindMessageService(message_code)

	if errMessage != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "A message data " + message_code + " are not found! Please try again!",
		})
		return
	}

	var replyMessageInput inputs.ReplyMessageInput

	errReplyMessageInput := c.ShouldBindJSON(&replyMessageInput)

	if errReplyMessageInput != nil {
		errorMessages := []string{}
		for _, e := range errReplyMessageInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	replyMessage, errReplyMessage := mc.messageService.ReplyMessageService(message_code, replyMessageInput, currentUser)

	if errReplyMessage != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when creating a message! Please try again!",
		})
		return
	}

	replyMessageRsps := responses.GetCreateMessageResponse(replyMessage)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new reply message for message code " + message_code + " is succesfully created!",
		"data":    replyMessageRsps,
	})

}

func (mc *messageController) UpdateMessageController(c *gin.Context) {
	message_code := c.Param("message_code")

	_, errMessage := mc.messageService.FindMessageService(message_code)

	if errMessage != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "A message data " + message_code + " are not found! Please try again!",
		})
		return
	}

	var updateMessageInput inputs.UpdateMessageInput

	errUpdateMessageInput := c.ShouldBindJSON(&updateMessageInput)

	if errUpdateMessageInput != nil {
		errorMessages := []string{}
		for _, e := range errUpdateMessageInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})
		return
	}

	currentUser := middleware.CurrentUser(c)

	updateUser, errUpdateUser := mc.messageService.UpdateMessageService(message_code, updateMessageInput, currentUser)

	if errUpdateUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when updating a message! Please try again!",
		})
		return
	}

	updateUserRsps := responses.GetUpdateMessageResponse(updateUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A message data from message code " + message_code + " is succesfully updated!",
		"data":    updateUserRsps,
	})
}

func (mc *messageController) DeleteMessageController(c *gin.Context) {
	message_code := c.Param("message_code")

	_, errMessage := mc.messageService.FindMessageService(message_code)

	if errMessage != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "400",
			"message": "A message data " + message_code + " are not found! Please try again!",
		})
		return
	}

	_, errDeleteMessage := mc.messageService.DeleteMessageService(message_code)

	if errDeleteMessage != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Error when deleting a message! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A message data from message code " + message_code + " is succesfully deleted!",
	})

}
