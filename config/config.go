package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var JWTKey string = ""

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
	jwtKey string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}

	JWTKey = app.jwtKey
	return &app
}

func EnvCloudName() string {
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		return val
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_CLOUD_NAME")
	}
}

func EnvCloudAPIKey() string {
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		return val
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_KEY")
	}
}

func EnvCloudAPISecret() string {
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		return val
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_API_SECRET")
	}
}

func EnvCloudUploadFolder() string {
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		return val
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	}
}
