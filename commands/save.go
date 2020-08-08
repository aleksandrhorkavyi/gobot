package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobot/models"
	"gobot/utils"
	"net/http"
)

type SaveCommand struct{
	Url string
}

func (c SaveCommand) Execute(u models.Update) (bool, error) {
	for len(u.Message.Entities) > 0 || utils.IsCallback(u) {
		return false, nil
	}
	botMessage := models.BotMessage{
		ChatId: u.Message.Chat.Id,
		Text: u.Message.User.FirstName + ", text was saved.",
	}
	buf, _ := json.Marshal(botMessage)
	fmt.Println(u.Message.Chat.Id)
	_, _  = http.Post(c.Url + "/sendMessage", "application/json", bytes.NewBuffer(buf))

	return true, nil
}
