package handlers

import (
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/bot"
	"context"
	"log"
)


func HandlerStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Println(update.Message.Chat.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "Спасибо за доверие сервису",
	})
}

func Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Println(update.Message.Chat.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: update.Message.Text,
	})
}

