package main

import (
	"userregister/infrastructure/config"

	_httpDeliver "userregister/delivery/http_handler"
	_useraccountRepo "userregister/repository"
	_useraccountUseCase "userregister/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	db := config.Init()

	route := gin.Default()

	//userregister repo and usecase
	userAccountRepo := _useraccountRepo.NewUserAccountRepository(db)
	userAccountUseCase := _useraccountUseCase.NewUserAccountUseCase(userAccountRepo)
	_httpDeliver.NewHttpHandler(route, userAccountUseCase)

	route.Run(viper.GetString("server.address"))

	defer db.Close()

}
