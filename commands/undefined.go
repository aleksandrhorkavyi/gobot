package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobot/models"
	"net/http"
)

type UndefinedCommand struct{
	Url string
}

func (c UndefinedCommand) Execute(u models.Update) (bool, error) {
	fmt.Println("Send not understand.")
	botMessage := models.BotMessage{
		ChatId: u.Message.Chat.Id,
		Text: u.Message.User.FirstName + ", I am not understand Your message. \n\nMaybe You combine some command with text for saving.",
	}

	buf, _ := json.Marshal(botMessage)
	fmt.Println(u.Message.Chat.Id)
	_, _  = http.Post(c.Url + "/sendMessage", "application/json", bytes.NewBuffer(buf))

	return true, nil
}
