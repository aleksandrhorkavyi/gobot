package models

type BotMessage struct {
	ChatId      int                   `json:"chat_id"`
	Text        string                `json:"text"`
	InlineReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type CallbackAnswer struct {
	CallbackQueryId string
	Text string
	ShowAlert bool
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	Id              string  `json:"id"`
	User            User    `json:"from"`
	Message         Message `json:"message"`
	ChatInstance    string  `json:"chat_instance"`
	Data            []byte  `json:"data"`
	InlineMessageId string  `json:"inline_message_id"`
}

type UserMessage struct {
	Text string
}

type Message struct {
	MessageId int      `json:"message_id"`
	Date      int      `json:"date"`
	Chat      Chat     `json:"chat"`
	User      User     `json:"from"`
	Text      string   `json:"text"`
	Photos    []Photo  `json:"photo"`
	Entities  []Entity `json:"entities"`
}

type Entity struct {
	Offset int
	Length int
	Type   string
}

func (e *Entity) IsCommand() bool {
	if e.Type == "bot_command" {
		return true
	}
	return false
}

type Chat struct {
	Id int `json:"id"`
}

type Photo struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	IsBot     bool   `json:"is_bot"`
}
