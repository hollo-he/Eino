package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 配置,变量读取啥的
type Config struct {
	OllamaModelName string
	OllamaUrl       string
}

var Cfg *Config

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg = &Config{
		OllamaModelName: os.Getenv("OLLAMA_MODEL_NAME"),
		OllamaUrl:       os.Getenv("OLLAMA_MODEL_URL"),
	}
}
