package setting

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var SecretId = ""

var SecretKey = ""

//var SecretId = "AKIDKSKwHL12nHUQ5TZPxFG"
//
//var SecretKey = "lFTQ6VeXuFJ5brqOffOq1"

var TencentApi = "dnspod.tencentcloudapi.com"

func init() {

	envFileName := ".env"
	_, err := os.Stat(envFileName)

	if os.IsNotExist(err) {
		file, err := os.Create(envFileName)
		if err != nil {
			log.Fatalf("Failed to create %s: %s", envFileName, err)
		}
		defer file.Close()

		_, err = file.WriteString("ID=A\nKEY=B")
		if err != nil {
			log.Fatalf("Failed to write to %s: %s", envFileName, err)
		}

		log.Printf("%s created with default values", envFileName)
	} else if err != nil {
		log.Fatalf("Failed to check %s: %s", envFileName, err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	SecretId = os.Getenv("ID")

	SecretKey = os.Getenv("KEY")
}
