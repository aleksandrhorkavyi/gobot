package callbacks

import (
	"fmt"
	"gobot/models"
	"gobot/utils"
)

type EndRestCallback struct {
	Url string
}

func (c EndRestCallback) Execute(u models.Update) (bool, error) {
	if !utils.IsCallback(u) {
		return false, nil
	}

	fmt.Println(u.Message.Chat.Id)

	utils.SendCallbackAnswer(&models.CallbackAnswer{
		CallbackQueryId: u.CallbackQuery.InlineMessageId,
		Text:   "Blet",
		ShowAlert: true,
	}, c.Url)

	return true, nil
}
