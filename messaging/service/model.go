package service

type BotMessage struct {
	MessageID int   `json:"message_id"`
	ChatID    int64 `json:"chat_id"`
	// From      *User  `json:"from"` // optional
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Date     int    `json:"date"`
	Text     string `json:"text"` // optional
}
