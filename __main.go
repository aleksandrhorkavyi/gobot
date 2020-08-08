package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const INTERVAL = 20
const REPEATS = 1

const botToken = "1214035824:AAEaO40XhSibhkFWxw2OAfOBd_vFrj5u5kA"
const botApi = "https://api.telegram.org/bot"
const botUrl = botApi + botToken

func main() {
	db, err := sql.Open("sqlite3", "./gobot.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scheduled_notes 
		(
		    id integer not null, 
		    chat_id integer not null, 
		    message_id integer not null, 
		    message text not null,
		    repeats integer not null,
			type	varchar(255) null
		)
	`)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user 
		(
		    chat_id integer not null, 
		    first_name varchar(255) null,
		    last_name varchar(255) null,
			username varchar(255) null
		)
	`)


	var offset int
	for {
		updates, err := getUpdates(offset)
		if err != nil {
			log.Println("err: ", err.Error())
		}
		for _, update := range updates {
			err = handle(update, db)
			offset = update.UpdateId + 1
			if err != nil {
				log.Println("err: ", err.Error())
			}
		}

		if time.Now().Minute() % INTERVAL != 0 {
			continue
		}

		if err := sendScheduledMessage(db); err != nil {
			log.Println("err: ", err.Error())
		}
	}
}

func sendScheduledMessage(db *sql.DB) error {
	notesRest := map[int]int{}
	scheduledNotesQuery, err := db.Query(`
		SELECT * FROM scheduled_notes
		WHERE repeats > 0
		GROUP BY chat_id
	`)

	if err != nil {
		return err
	}
	for scheduledNotesQuery.Next() {
		var message DbMessage
		scheduledNotesQuery.Scan(
			&message.Id,
			&message.ChatId,
			&message.MessageId,
			&message.Message,
			&message.Repeats,
			&message.Type,
		)
		user, err := getDbUserByChatId(message.ChatId, db)
		if err != nil {
			return err
		}
		notesRest[message.Id] = message.Repeats - 1
		var botMessage BotMessage
		botMessage.ChatId = message.ChatId
		botMessage.Text = message.Message

		buf, err := json.Marshal(botMessage)
		if err != nil {
			return err
		}

		if message.Type == "photo" {
			_, _ = getBotForwardMessage(message)
			fmt.Println("Forward 1 photo to " + user.FirstName)
		}

		if message.Type == "text" {
			_, _  = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
			fmt.Println("Send 1 message to " + user.FirstName)
		}
	}
	scheduledNotesQuery.Close()
	for id, repeats := range notesRest {
		_, err := db.Exec("UPDATE scheduled_notes SET repeats = $1 WHERE id = $2", repeats, id)
		if err != nil {
			return nil;
		}
	}

	return nil
}

func getDbUserByChatId(id int, db *sql.DB) (DbUser, error) {
	var user DbUser
	stmt, err := db.Query("SELECT * FROM user WHERE chat_id = $1 LIMIT 1", id)
	defer stmt.Close()
	if err != nil {
		return user, err
	}

	for stmt.Next() {
		stmt.Scan(&user.ChatId, &user.FirstName, &user.LastName, &user.Username)
	}

	return user, nil
}

func getUpdates(offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func handle(update Update, db *sql.DB) error {
	fmt.Println("Received 1 message from " + update.Message.User.FirstName)
	if err := saveUser(update, db); err != nil {
		return err
	}
	notesStmt, _ := db.Prepare("INSERT INTO scheduled_notes (id, chat_id, message_id, message, repeats, type) VALUES (?, ?, ?, ?, ?, ?)")

	var msgType string
	if len(update.Message.Photos) > 0 {
		msgType = "photo"
	} else {
		msgType = "text"
	}
	_, _ = notesStmt.Exec(update.UpdateId, update.Message.Chat.Id, update.Message.MessageId, update.Message.Text, REPEATS, msgType)

	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.Id
	botMessage.Text = update.Message.User.FirstName + ", your " + msgType +" was saved."
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err  = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func getBotForwardMessage(dbMessage DbMessage) (BotForwardMessage, error) {
	var message BotForwardMessage
	message.ChatId = dbMessage.ChatId
	message.FromChatId = dbMessage.ChatId
	message.MessageId = dbMessage.MessageId

	buf, err := json.Marshal(message)
	if err != nil {
		return message, err
	}
	_, err = http.Post(botUrl+"/forwardMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return message, err
	}

	return message, nil
}

func saveUser(update Update, db *sql.DB) error {
	userQuery, err := db.Query("SELECT * FROM user WHERE chat_id = $1 LIMIT 1", update.Message.Chat.Id)
	if err != nil {
		return err
	}

	defer userQuery.Close()
	for userQuery.Next() {
		return nil
	}

	userStmt, _ := db.Prepare("INSERT INTO user (chat_id, first_name, last_name, username) VALUES (?, ?, ?, ?)")
	_, userErr := userStmt.Exec(
		update.Message.Chat.Id,
		update.Message.User.FirstName,
		update.Message.User.LastName,
		update.Message.User.Username,
	)

	if userErr != nil {
		return err
	}

	return nil
}