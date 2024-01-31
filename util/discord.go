package util

import (
	"github.com/gtuk/discordwebhook"
	"github.com/labstack/gommon/log"
	"github.com/minnowo/astoryofand/assets"
)

func SendDiscordOrderWebhook(content string) {

	if len(assets.DiscordWebhookURL) < 10 {
		log.Error("Discord webhook url is invalid")
		return
	}

	if len(content) == 0 {
		log.Warn("Tried to send empty Discord webhook message")
		return
	}

	message := discordwebhook.Message{
		Username: &assets.DiscordWebhookName,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(assets.DiscordWebhookURL, message)

	log.Debug("Sent Discord webhook")

	if err != nil {
		log.Error(err)
	}
}
