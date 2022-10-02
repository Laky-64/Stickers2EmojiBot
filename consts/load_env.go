package consts

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	BotToken = os.Getenv("BOT_TOKEN")
}
