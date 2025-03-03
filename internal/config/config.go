package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	path := ".env"
	for {
		err := godotenv.Load(path)
		if err == nil {
			break
		}
		path = "../" + path
	}
}
