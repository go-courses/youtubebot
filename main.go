package main

import (
	"flag"
	"log"
	"os"
	"youtubebot/bot"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	// используя токен создаем новый инстанс бота
	newBot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", newBot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := newBot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "sound" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Какой звук вы хотите?")
			newBot.Send(msg)
		}

		text := update.Message.Text

		id := bot.Search(text)

		// создание сообщения от бота пользвателю
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, id)
		newBot.Send(msg)

	}
}
