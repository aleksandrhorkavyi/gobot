package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobot/models"
	"net/http"
)

func SendMessage(msg *models.BotMessage, url string) *http.Response {
	buf, _ := json.Marshal(msg)
	resp, err := http.Post(
		url+"/sendMessage",
		"application/json",
		bytes.NewBuffer(buf),
	)
	if err != nil {
		panic(err)
	}

	return resp
}

func SendCallbackAnswer(answer *models.CallbackAnswer, url string) *http.Response {
	buf, _ := json.Marshal(answer)
	resp, err := http.Post(
		url+"/answerCallbackQuery",
		"application/json",
		bytes.NewBuffer(buf),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	return resp
}
