package bot

func Start() {

	bot, _ := CreateBot(telegramBotToken)
	updates, _ := CreateChannel(bot)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		text := update.Message.Text
		id := Search(text)
		SendMsg(update, bot, "Начал поиск")
		url := "https://www.youtube.com/watch?v=" + id
		SendMsg(update, bot, url)
		//url, _ := GetDownloadUrl(id)

	}
}
