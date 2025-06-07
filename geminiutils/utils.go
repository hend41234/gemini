package geminiutils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Utilization struct {
	GeminiApiKey string
	BaseURL      string
	Endpoint     map[string]string
}

var Utils *Utilization

func LoadConf(envFile string) {
	if loadErr := godotenv.Load(envFile); loadErr != nil {
		log.Fatal("env file not found")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY not found, makse sure the key of API KEY is 'GEMINI_API_KEY'")
	}
}

func init() {
	if loadErr := godotenv.Load(".env"); loadErr != nil {
		return
	}
	if apiKey := os.Getenv("GEMINI_API_KEY"); apiKey == "" {
		fmt.Println("your config unset!\nfor default name env file is .env, but you can set according you want with use LoadConf('envfile')")
		return
	}

	Utils = new(Utilization)
	Utils.GeminiApiKey = os.Getenv("GEMINI_API_KEY")
	Utils.BaseURL = "https://generativelanguage.googleapis.com/v1beta/models"
	Utils.Endpoint = map[string]string{
		"2FL":   "gemini-2.0-flash-lite:generateContent?key=",
		"2FLS":  "gemini-2.0-flash-lite:streamGenerateContent?alt=sse&key=",
		"2F":    "gemini-2.0-flash:generateContent?key=",
		"2.5FL": "gemini-2.5-flash-lite:generateContent?key=",
	}
}
