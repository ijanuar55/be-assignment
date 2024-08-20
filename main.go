package main

import (
	"be-assignment/config"
	"be-assignment/controller"
	"be-assignment/helper"
	"be-assignment/repository"
	"be-assignment/routes"
	"be-assignment/service"
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

	// repository
	userRepository := repository.NewUserRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	// service
	userService := service.NewUserServiceImpl(userRepository)
	accountService := service.NewAccountServiceImpl(accountRepository)
	transactionService := service.NewTransactionServiceImpl(transactionRepository, accountRepository)

	// controller
	userController := controller.NewUserController(userService)
	accountController := controller.NewAccountController(accountService, userService)
	transactionController := controller.NewTransactionController(transactionService, accountService)

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

	r := routes.Router(userController, accountController, transactionController)

	r.Run()
}
