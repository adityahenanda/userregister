package http_handler

import (
	"userregister/middleware"
	"userregister/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var jwtService middleware.JWTService = middleware.NewJwtService()

type HandlerUserAccount struct {
	UserAccountUseCase models.UserAccountUsecase
	JwtService         middleware.JWTService
}

func NewHttpHandler(route *gin.Engine, userAccountUseCase models.UserAccountUsecase) {

	handlerUserAccount := &HandlerUserAccount{
		UserAccountUseCase: userAccountUseCase,
		JwtService:         jwtService,
	}

	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		{
			useraccount := api.Group("/user")
			{
				useraccount.POST("/register", handlerUserAccount.RegisterUserAccountHandler)
				useraccount.POST("/login", handlerUserAccount.LoginUserAccountHandler)
			}

			useraccountAuth := api.Group("/auth", middleware.AuthorizeJWT(jwtService))
			{
				useraccountAuth.GET("/user-account/:userId", handlerUserAccount.GetUserAccountById)
				useraccountAuth.GET("/user-account", handlerUserAccount.GetAllUserAccount)
				useraccountAuth.PUT("/user-account/:userId", handlerUserAccount.UpdateUserAccount)
				useraccountAuth.DELETE("/user-account/:userId", handlerUserAccount.DeleteUserAccount)
			}
		}
	}
}
