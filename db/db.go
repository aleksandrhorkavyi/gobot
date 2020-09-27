package db

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
	"gobot/config"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Deploy(c *cli.Context, config *config.Config) bool {
	db, err := sql.Open("sqlite3", "./gobot.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(config.Db.MigrationsDirectory)
	if err != nil {
		panic(err)
	}

	fmt.Println("New migrations:")
	for _, f := range files {
		fmt.Println("\t- " + f.Name())
	}

	if agreeded := answerUserAgree(); !agreeded{
		return false
	}

	for _, f := range files {
		contentBytes, err := ioutil.ReadFile(config.Db.MigrationsDirectory + f.Name())
		if err != nil {
			panic(err)
		}
		fmt.Println("Run migration: " + f.Name())
		res, err := db.Exec(string(contentBytes))
		if err != nil {
			panic(err)
		}
		fmt.Println("Migration passed")
		fmt.Println(res.RowsAffected())
	}

	return true
}

func answerUserAgree() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are You sure to run this migrations? [y/n]: ")
	text, _ := reader.ReadString('\n')

	if strings.Trim(text, "\n") == "y" {
		return true
	} else if strings.Trim(text, "\n") == "n" {
		return false
	}
	fmt.Println("Please, type y or n")
	return answerUserAgree()
}

func Connection() *sql.DB {
	db, err := sql.Open("sqlite3", "./gobot.db")
	if err != nil {
		panic(err)
	}
	return db
}
