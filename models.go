package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int     `json:"message_id"`
	Date      int     `json:"date"`
	Chat      Chat    `json:"chat"`
	User      User    `json:"from"`
	Text      string  `json:"text"`
	Photos    []Photo `json:"photo"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	IsBot     bool   `json:"is_bot"`
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

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type BotForwardMessage struct {
	ChatId     int `json:"chat_id"`
	FromChatId int `json:"from_chat_id"`
	MessageId  int `json:"message_id"`
}

type DbMessage struct {
	Id      int
	ChatId  int
	MessageId int
	Message string
	Repeats int
	Type    string
}

type DbUser struct {
	ChatId    int
	FirstName string
	LastName  string
	Username  string
}
