package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"github.com/go-telegram/bot"
	"tele/tools/env_working"
	"tele/tools/handlers"
)



func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	bot_key := env_working.GettingEnv()
	defer cancel()
	opts := []bot.Option{
		bot.WithCheckInitTimeout(1 * time.Minute),
		bot.WithDefaultHandler(handlers.Handler),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, handlers.HandlerStart),
	}
	b, err := bot.New(bot_key, opts...)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	b.Start(ctx)
}

