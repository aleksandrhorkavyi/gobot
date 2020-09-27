package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"gobot/config"
	"gobot/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Run(conf *config.Config) {

	api := &BotApi{url: conf.Bot.Url}

	for {
		updates, err := getUpdates(api)
		if err != nil {
			panic(err.Error())
		}
		for _, update := range updates {
			i := &Invoker{
				update: update,
				api:    api,
			}
			registerActions(i, conf)
			i.ExecuteActions()
		}
	}
}

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

type BotApi struct {
	url    string
	offset int
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
