package commands

import (
	"gobot/models"
	"gobot/utils"
)

type RestCommand struct {
	Url string
}

func (c RestCommand) Execute(u models.Update) (bool, error) {

	if !utils.IsCommand(u, "/rest") {
		return false, nil
	}

	timePicker := models.TimePicker{Type: "rest_timepicker"}

	scheduleInlineKeyboard := [][]models.InlineKeyboardButton{
		{
			models.InlineKeyboardButton{Text: "00.00", CallbackData: timePicker.As("s0")},
			models.InlineKeyboardButton{Text: "01.00", CallbackData: timePicker.As("s1")},
			models.InlineKeyboardButton{Text: "02.00", CallbackData: timePicker.As("s2")},
			models.InlineKeyboardButton{Text: "03.00", CallbackData: timePicker.As("s3")},
		},
		//{
		//	models.InlineKeyboardButton{Text: "04.00", CallbackData: timePicker.As("04:00")},
		//	models.InlineKeyboardButton{Text: "05.00", CallbackData: timePicker.As("05:00")},
		//	models.InlineKeyboardButton{Text: "06.00", CallbackData: timePicker.As("06:00")},
		//	models.InlineKeyboardButton{Text: "07.00", CallbackData: timePicker.As("07:00")},
		//},
		//{
		//	models.InlineKeyboardButton{Text: "08.00", CallbackData: timePicker.As("08.00")},
		//	models.InlineKeyboardButton{Text: "09.00", CallbackData: timePicker.As("09.00")},
		//	models.InlineKeyboardButton{Text: "10.00", CallbackData: timePicker.As("10.00")},
		//	models.InlineKeyboardButton{Text: "11.00", CallbackData: timePicker.As("11.00")},
		//},
		//{
		//	models.InlineKeyboardButton{Text: "12.00", CallbackData: timePicker.As("12.00")},
		//	models.InlineKeyboardButton{Text: "13.00", CallbackData: timePicker.As("13.00")},
		//	models.InlineKeyboardButton{Text: "14.00", CallbackData: timePicker.As("14.00")},
		//	models.InlineKeyboardButton{Text: "15.00", CallbackData: timePicker.As("15.00")},
		//},
		//{
		//	models.InlineKeyboardButton{Text: "16.00", CallbackData: timePicker.As("16.00")},
		//	models.InlineKeyboardButton{Text: "17.00", CallbackData: timePicker.As("17.00")},
		//	models.InlineKeyboardButton{Text: "18.00", CallbackData: timePicker.As("18.00")},
		//	models.InlineKeyboardButton{Text: "19.00", CallbackData: timePicker.As("19.00")},
		//},
		//{
		//	models.InlineKeyboardButton{Text: "20.00", CallbackData: timePicker.As("20.00")},
		//	models.InlineKeyboardButton{Text: "21.00", CallbackData: timePicker.As("21.00")},
		//	models.InlineKeyboardButton{Text: "22.00", CallbackData: timePicker.As("22.00")},
		//	models.InlineKeyboardButton{Text: "23.00", CallbackData: timePicker.As("23.00")},
		//},
	}

	utils.SendMessage(&models.BotMessage{
		ChatId: u.Message.Chat.Id,
		//Text:   "Ок, с какого часа вы не хотите ничего делать?",
		Text:   "GO!",
		InlineReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: scheduleInlineKeyboard,
		},
	}, c.Url)

	return true, nil
}
