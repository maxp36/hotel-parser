package parser

// Parser represent the parser's service
type Parser interface {
	Parse(data []byte) error

	// Message
	// SendMessage(userID, partnerID uint64, text string, objects []models.MessageObject) (userChat *models.Chat, partnerChat *models.Chat, err error)
	// DeleteMessage(userID, mesID uint64) (userChat *models.Chat, err error)
}
