package bot

func Start() {
	const telegramBotToken = "DEVELOPER KEY PASTE HERE"

	bot, _ := CreateBot(telegramBotToken)
	updates, _ := CreateChannel(bot)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "sound" {
			SendMsg(update, bot, "Какой звук вы хотите найти?")
		}

		text := update.Message.Text
		id := Search(text)

		SendMsg(update, bot, id)

	}
}
