package bot

import (
	"log"
	"os/exec"
)

const (
	telegramBotToken = ""
	developerKey     = ""
)

func Start() {

	bot, _ := CreateBot(telegramBotToken)
	updates, _ := CreateChannel(bot)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		// получаем любой текст от пользователя
		text := update.Message.Text

		// используем Youtube API чтобы найти самый подходящий запрос под этот текст
		id := Search(text)

		// чтобы пользователь понимал, что бот работает, выводим в чат сообщение
		SendMsg(update, bot, "Начал поиск")

		// пока что чат бот выдаёт найденное видео в виде полной ссылки и отправляет его
		url, title, _ := GetDownloadUrl(id)

		// скачивает видео и конвертирует его в mp3 например
		Convert(title, url)

		// для пользователя, чтобы знал, что бот работает.
		SendMsg(update, bot, "Начал конвертацию")

		// создает линк, чтобы закидывать файл в
		link := "files/" + title + ".mp3"

		// в конце загружается готовое аудио через указанный путь.
		SendAudio(update, bot, link)

		cmd := exec.Command("rm", link)
		err := cmd.Run()
		if err != nil {
			log.Printf("Ошибка %s", err)
		}
	}
}
