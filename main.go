package main

import (
	"flag"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"chatbotproject/telegram"
	"chatbotproject/youtube"
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
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "sound" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Какой звук вы хотите?")
			bot.Send(msg)
		}

		text := update.Message.Text
		
		// создание сообщения от бота пользвателю
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)

		
		
		
	}
}