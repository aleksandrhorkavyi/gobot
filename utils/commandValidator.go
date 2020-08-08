package utils

import (
	//"encoding/json"
	"gobot/models"
)

func IsCommand(u models.Update, command string) (bool) {
	for _, entity := range u.Message.Entities {
		if !entity.IsCommand() {
			return false
		}
	}
	if u.Message.Text != command {
		return false
	}
	return true
}

func IsCallback(u models.Update) (bool) {
	//var timePicker models.TimePicker
	//err := json.Unmarshal(u.CallbackQuery.Data, &timePicker)
	if u.CallbackQuery.Data != nil {
		return true
	}
	return false
}
