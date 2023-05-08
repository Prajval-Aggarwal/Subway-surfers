package server

import (
	_ "subway/docs"
	"subway/server/gateway"
	"subway/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//starting cron
	//handler.StartCron()

	//TODO
	//Apply middleware to required routes.. not added for testing thing

	//authentication route
	server.engine.POST("/register-player", handler.RegisterHandler)
	server.engine.POST("/login", handler.LoginHandler)
	server.engine.DELETE("/logout", gateway.PlayerAuthentication, handler.LogoutHandler)
	server.engine.PATCH("/update-pass", gateway.PlayerAuthentication, handler.UpdatePasswordHandler)
	server.engine.PATCH("/update-name", gateway.PlayerAuthentication, handler.UpdateNameHandler)
	server.engine.POST("/forgot-password", handler.ForgotPasswordHandler)
	server.engine.PATCH("/reset-password", handler.ResetPasswordHandler)

	//player detail route
	server.engine.GET("/show-player", handler.ShowPlayerDetailsHandler)

	//powerup routes
	server.engine.GET("/show-powerups", handler.ShowPowerUpsHandler)
	server.engine.POST("/use-powerup", gateway.PlayerAuthentication, handler.UsePowerUpHandler)
	server.engine.POST("/buy-powerup", gateway.PlayerAuthentication, handler.BuyPowerupHandler)

	//reward handler
	server.engine.GET("/collect-reward", gateway.PlayerAuthentication, handler.RewardCollectedHandler)
	server.engine.GET("/show-reward", gateway.PlayerAuthentication, handler.ShowPlayerRewardHandler)

	//leaderboard route
	server.engine.GET("/show-leaderboard", handler.ShowLeaderBoardHandler)

	//end game route
	server.engine.POST("/end-game", gateway.PlayerAuthentication, handler.EndGameHandler)

	//payment route
	server.engine.POST("/make-payment", gateway.PlayerAuthentication, handler.MakePaymentHandler)

	//cart routes
	server.engine.GET("/show-cart", handler.ShowCartHandler)

	//avatar route
	server.engine.GET("/show-avatars", handler.ShowAvatarHandler)
	server.engine.PATCH("/update-avatar", gateway.PlayerAuthentication, handler.UpdateAvatarHandler)

	//swaggger route
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
