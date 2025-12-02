package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 配置,变量读取啥的
type Config struct {
	OllamaToolModelName string
	OllamaChatModelName string
	OllamaUrl           string
	ZhipuAiName         string
	ZhipuAiUrl          string
	ZhipuApiKey         string
}

var Cfg *Config

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg = &Config{
		OllamaToolModelName: os.Getenv("OLLAMA_MODEL_NAME_TOOL"),
		OllamaUrl:           os.Getenv("OLLAMA_MODEL_URL"),
		OllamaChatModelName: os.Getenv("OLLAMA_MODEL_NAME_CHAT"),
		ZhipuAiName:         os.Getenv("ZHIPU_AI_MODEL"),
		ZhipuAiUrl:          os.Getenv("ZHIPU_AI_BASE_URL"),
		ZhipuApiKey:         os.Getenv("ZAI_API_KEY"),
	}
}
