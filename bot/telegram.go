package bot

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// создает бота
func CreateBot(telegramBotToken string) (*tgbotapi.BotAPI, error) {
	// используя токен создаем новый инстанс бота
	newBot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", newBot.Self.UserName)

	return newBot, err
}

//создаёт канал
func CreateChannel(newBot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := newBot.GetUpdatesChan(u)
	return updates, err
}

// отправляет заданное сообщение в канал заданного бота
func SendMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI, msg string) {
	text := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	bot.Send(text)
}

// отправляет аудиофайл в канал заданного бота

func SendAudio(update tgbotapi.Update, bot *tgbotapi.BotAPI, filePath string) {
	audio := tgbotapi.NewAudioUpload(update.Message.Chat.ID, filePath)
	bot.Send(audio)
}
