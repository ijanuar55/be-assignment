package main

import (
	"be-assignment/config"
	"be-assignment/helper"
	"be-assignment/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Handle DB Connection
	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	// initialize supertokens
	err = supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("CONNECTION_URI"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:       os.Getenv("APP_NAME"),
			APIDomain:     os.Getenv("DOMAIN"),
			WebsiteDomain: os.Getenv("DOMAIN"),
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})

	if err != nil {
		log.Fatalf("Could not initialize SuperTokens: %v", err)
	}

	r := routes.Router()

	r.Run()
}
