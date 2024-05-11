package swagger

import (
	"log"
	"os"
)

func GetApiServerAddress() string {

	address, exists := os.LookupEnv("API_SERVER_ADDRESS")
	if !exists {
		log.Fatal("ENV variable API_SERVER_ADDRESS not set. Forgott to load a .env file?")
	}

	return address

}

func GetDbmsAddress() string {

	address, exists := os.LookupEnv("DBMS_ADDRESS")
	if !exists {
		log.Fatal("ENV variable DBMS_ADDRESS not set. Forgott to load a .env file?")
	}

	return address

}
