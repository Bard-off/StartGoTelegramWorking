package env_working

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)



func GettingEnv() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	bot_key := os.Getenv("BOT_KEY")
	log.Println("Получена информация с .env файла")
	return bot_key
}