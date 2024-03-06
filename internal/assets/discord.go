package assets

import "os"

var DiscordWebhookName string = "A Story Of And"

var DiscordWebhookURL string = os.Getenv("DISCORD_WEBHOOK")
