package config

import (
	"log"
	"os"
)

var (
	AccessToken = ""
	Port        = "8060"
)

func init() {
	if accessToken, found := os.LookupEnv(EnvNameAccessToken); found {
		AccessToken = accessToken
	}
	if port, found := os.LookupEnv(EnvNameServerPort); found {
		Port = port
	}
	if AccessToken == "" {
		log.Fatalln("ACCESS_TOKEN is empty.")
	}
}
