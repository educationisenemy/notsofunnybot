package main

import (
	"github.com/educationisenemy/notsofunnybot/news/kaktam"
	"github.com/Syfaro/telegram-bot-api"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type appConfig struct {
	Token string
}

func readAppConfig() (appConfig, error) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config := appConfig{}
	err := decoder.Decode(&config)

	return config, err
}

func makeTextFromLink(link kaktam.Link) string {
	text := link.Text + "\n" + link.Href
	return text
}

func main() {

	config, err := readAppConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(ucfg)

	for {
		select {
		case update := <-updates:
			UserName := update.Message.From.UserName
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			if strings.Contains(Text, "/") {
				links := kaktam.GrabLinks(99999);

				for index, element := range links {
					_ = index
					msg := tgbotapi.NewMessage(ChatID, makeTextFromLink(element))
					bot.Send(msg)
				}
			}
		}
	}
}