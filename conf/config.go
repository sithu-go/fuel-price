package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load env file
	currentDir, _ := os.Getwd()
	log.Println(currentDir + "/conf/.env")
	err := godotenv.Load(currentDir + "/conf/.env")
	if err != nil {
		log.Println("error opening .env file")
		log.Fatalf(err.Error(), "FGDD")
		return
	}
}
