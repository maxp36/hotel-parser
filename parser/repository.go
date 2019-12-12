package parser

// Repository represent the parser's Postgres repository contract.
type Repository interface {
	AddHotel() error
	AddHotelImage() error

	// AddChatIfDoesNotExist(userID, partnerID uint64) (*models.Chat, error)
	// GetChatIDByUserIDAndPartnerID(userID, partnerID uint64) (uint64, error)
	// GetChatByID(chatID, mLimit, mOffset uint64) (*models.Chat, error)
	// GetChatsByUserID(userID, limit, offset uint64) ([]models.Chat, error)
	// DeleteChat(chatID uint64) (*models.Chat, error)
	// ClearChat(chatID uint64) (*models.Chat, error)
}
