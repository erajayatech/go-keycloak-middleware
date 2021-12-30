package keycloakmiddleware

import (
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

func getEnvOrDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load file .env: ", err)
		panic(err)
	}

	value := getEnvOrDefault(key, "").(string)
	return value
}

func decodeBase64BigInt(s string) *big.Int {
	buffer, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	return big.NewInt(0).SetBytes(buffer)
}
