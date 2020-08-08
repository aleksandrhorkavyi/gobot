package commands

import (
	"gobot/models"
	"gobot/utils"
)

type MenuCommand struct {
	Url string
}

func (c MenuCommand) Execute(u models.Update) (bool, error) {
	if !utils.IsCommand(u, "/menu") {
		return false, nil
	}

	utils.SendMessage(&models.BotMessage{
		ChatId: u.Message.Chat.Id,
		Text:   "Your menu",
	}, c.Url)

	return true, nil
}
