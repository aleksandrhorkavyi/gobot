package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gobot/callbacks"
	"gobot/commands"
	"gobot/config"
	"gobot/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BotApi struct {
	url    string
	offset int
}

func registerActions(i *Invoker, conf *config.Config) {
	i.Register(commands.RestCommand{Url: conf.BotUrl()})
	i.Register(commands.MenuCommand{Url: conf.BotUrl()})
	i.Register(commands.SaveCommand{Url: conf.BotUrl()})
	i.Register(callbacks.EndRestCallback{Url: conf.BotUrl()})

	i.RegisterUndefined(commands.UndefinedCommand{Url: conf.BotUrl()})
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	http.HandleFunc("/gobot", handleWebHook) // Устанавливаем роутер
	err := http.ListenAndServe(":3000", nil) // устанавливаем порт веб-сервера
	fmt.Println("Start blet")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebHook(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &models.Update{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "marco") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sayPolo(body.Message.Chat.Id); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

func sayPolo(chatID int) error {
	// Create the request body struct
	reqBody := &BotMessage{
		ChatId: chatID,
		Text:   "Polo!!",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot1214035824:AAEaO40XhSibhkFWxw2OAfOBd_vFrj5u5kA/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

//func main() {
//	conf := config.New()
//	api := &BotApi{url: conf.Bot.Url}
//
//	for {
//		updates, err := getUpdates(api)
//		if err != nil {
//			panic(err.Error())
//		}
//		for _, update := range updates {
//			i := &Invoker{
//				update: update,
//				api:    api,
//			}
//			registerActions(i, conf)
//			i.ExecuteActions()
//		}
//	}
//}

func getUpdates(api *BotApi) ([]models.Update, error) {
	resp, err := http.Get(api.url + "/getUpdates" + "?offset=" + strconv.Itoa(api.offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse models.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}
