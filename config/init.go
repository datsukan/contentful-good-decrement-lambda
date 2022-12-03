package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "contentful-good-decrement-lambda"

// LoadContentfulEnv は Contentful SDK の接続情報を環境変数から読み込む
func LoadContentfulEnv() (string, string, error) {
	env := os.Getenv("ENV")

	if env != "production" {
		projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		currentWorkDirectory, _ := os.Getwd()
		rootPath := projectName.Find([]byte(currentWorkDirectory))
		var envFilePath string
		if string(rootPath) == "" {
			envFilePath = "./.env"
		} else {
			envFilePath = string(rootPath) + `/.env`
		}

		err := godotenv.Load(envFilePath)
		if err != nil {
			return "", "", err
		}
	}

	token := os.Getenv("CONTENTFUL_ACCESS_TOKEN")
	spaceID := os.Getenv("CONTENTFUL_SPACE_ID")

	if token == "" || spaceID == "" {
		m := fmt.Sprintf("Environment variable not set. [ token: %v, spaceID: %v ]", token, spaceID)
		fmt.Println(m)
		return "", "", errors.New(m)
	}

	return token, spaceID, nil
}
