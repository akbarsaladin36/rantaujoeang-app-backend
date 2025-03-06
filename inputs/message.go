package inputs

type CreateMessageInput struct {
	ReceiverId         string `json:"receiver_id"`
	MessageTitle       string `json:"message_title"`
	MessageDescription string `json:"message_description"`
}

type ReplyMessageInput struct {
	MessageTitle       string `json:"message_title"`
	MessageDescription string `json:"message_description"`
}

type UpdateMessageInput struct {
	MessageTitle       string `json:"message_title"`
	MessageDescription string `json:"message_description"`
}
