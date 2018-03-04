package bot

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-courses/youtubebot/conf"
	"github.com/pkg/errors"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	telegramAPIUpdateInterval = 60
)

// Bot ...
type Bot struct {
	c       conf.BotConfig
	tgAPI   *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	yClient *youtube.Service
}

// NewTGBot creates a new bot
func NewTGBot(c conf.BotConfig) (*Bot, error) {
	newBot, err := tgbotapi.NewBotAPI(c.TelegramToken)
	if err != nil {
		return nil, errors.Wrap(err, "could not create bot")
	}
	b := &Bot{
		c:     c,
		tgAPI: newBot,
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = telegramAPIUpdateInterval
	updates, err := b.tgAPI.GetUpdatesChan(u)
	if err != nil {
		return nil, errors.Wrap(err, "could not create updates chan")
	}
	b.updates = updates
	client, err := youtube.New(&http.Client{
		Transport: &transport.APIKey{Key: c.YoutubeDeveloperKey},
	})
	if err != nil {
		return nil, err
	}
	b.yClient = client
	return b, nil
}

func (b *Bot) sendMsg(update tgbotapi.Update, msg string) {
	text := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	b.tgAPI.Send(text)
}

func (b *Bot) sendAudio(update tgbotapi.Update, filePath string) {
	audio := tgbotapi.NewAudioUpload(update.Message.Chat.ID, filePath)
	b.tgAPI.Send(audio)
}

// Start ...
func (b *Bot) Start() {
	fmt.Println("Starting tg bot")
	for update := range b.updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		if text == "" {
			continue
		}
		youtubeID, err := b.search(text)
		if err != nil {
			log.Println("could not get video id from youtube", err)
		}
		b.sendMsg(update, "Начал поиск")
		url, title, err := GetDownloadURL(youtubeID)
		if err != nil {
			log.Println("could not get download url", err)
		}
		err = Convert(title, url)
		if err != nil {
			log.Println("could not convert video file to mp3 ", err)
		}
		b.sendMsg(update, "Начал конвертацию")
		link := "files/" + title + ".mp3"
		b.sendAudio(update, link)
		os.Remove(link)

	}
}
