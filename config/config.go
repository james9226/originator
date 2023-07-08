package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	Env                     string
	ApiToken                string
	ProjectID               string
	PrivateKeyID            string
	PrivateKey              string
	ClientEmail             string
	ClientID                string
	AuthURI                 string
	TokenURI                string
	AuthProviderX509CertURL string
	ClientX509CertURL       string
}

var cfg *Config

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg = &Config{
		Env:                     os.Getenv("ENV"),
		ApiToken:                os.Getenv("API_TOKEN"),
		ProjectID:               os.Getenv("PROJECT_ID"),
		PrivateKeyID:            os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("PRIVATE_KEY"),
		ClientEmail:             os.Getenv("CLIENT_EMAIL"),
		ClientID:                os.Getenv("CLIENT_ID"),
		AuthURI:                 os.Getenv("AUTH_URI"),
		TokenURI:                os.Getenv("TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       os.Getenv("CLIENT_X509_CERT_URL"),
	}

	// Iterate over the fields of the Config struct
	val := reflect.ValueOf(*cfg)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			return fmt.Errorf("Missing necessary environment variables")
		}
	}

	return nil
}

func Get() *Config {
	return cfg
}
