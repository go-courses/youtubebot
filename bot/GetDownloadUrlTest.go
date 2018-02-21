package main

import (
	"log"
	"net/url"

	"github.com/rylio/ytdl" // Чтобы установить go get -u github.com/rylio/ytdl/
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	
bot, err := tgbotapi.NewBotAPI("Здесь_Token_бота")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		url, _ := GetDownloadUrl("Здесь_id_видео")
		urlString := url.String() // Если нужна ссылка в виде string 
		sendMessage := tgbotapi.NewMessage(update.Message.Chat.ID, urlString)
		bot.Send(sendMessage)
	}
}

func GetDownloadUrl(idVideo string) (*url.URL, error) {

	infoFromId, _ := ytdl.GetVideoInfoFromID(idVideo)

	foundFormat := func (formats ytdl.FormatList) ytdl.Format {
		var foundFormat ytdl.Format
		bestFormats := formats.Filter(ytdl.FormatResolutionKey, []interface{}{/*Разрешение искомого видео*/"360p", "720p"}).Filter(ytdl.FormatExtensionKey, []interface{}{/*Расширение видео*/"mp4"}).Filter(ytdl.FormatAudioEncodingKey, []interface{}{/*Формат звука*/"aac"}).Extremes(ytdl.FormatResolutionKey, true).Extremes(ytdl.FormatAudioBitrateKey, true)
		for _, format := range bestFormats {
			if format.Extension == "mp4" {
				foundFormat = format
			}
		}
		return foundFormat
	}
	downloadUrl, err := infoFromId.GetDownloadURL(foundFormat(infoFromId.Formats))
	return downloadUrl, err
}