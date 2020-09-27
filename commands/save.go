package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/doug-martin/goqu"
	"gobot/db"
	"gobot/models"
	"gobot/utils"
	"log"
	"net/http"
)

type SaveCommand struct{
	Url string
}

func (c SaveCommand) Execute(u models.Update) (bool, error) {
	for len(u.Message.Entities) > 0 || utils.IsCallback(u) {
		return false, nil
	}

	sql, _, _ := goqu.Insert("scheduled_notes").Rows(
			goqu.Record{
				"id": u.Message.MessageId,
				"chat_id": u.Message.Chat.Id,
				"message_id": u.Message.MessageId,
				"message": u.Message.Text,
				"repeats": 1,
				"type": "text",
			},
		).ToSQL()

	conn := db.Connection()
	defer conn.Close()
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal(err)
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
