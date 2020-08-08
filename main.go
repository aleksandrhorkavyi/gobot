package main

import (
	"encoding/json"
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
	http.HandleFunc("/", sayhello) // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//func main() {
//	http.HandleFunc("/", sayhello) // Устанавливаем роутер
//	err := http.ListenAndServeTLS(
//		":8080",
//		"/Users/aleksandrhorkavyi/go/src/gobot/cert.pem",
//		"/Users/aleksandrhorkavyi/go/src/gobot/key.pem",
//		nil,
//	)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет Нах!")
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
