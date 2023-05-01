package server

import (
	_ "subway/docs"
	"subway/server/handler"
	"subway/server/provider"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//authentication route
	server.engine.POST("/register-player", handler.RegisterHandler)
	server.engine.POST("/login", handler.LoginHandler)
	server.engine.DELETE("/logout", handler.LogoutHandler)
	server.engine.PATCH("/update-pass", provider.PlayerAuthentication, handler.UpdatePasswordHandler)
	server.engine.PATCH("/update-name", provider.PlayerAuthentication, handler.UpdateNameHandler)
	server.engine.POST("/forgot-password", handler.ForgotPasswordHandler)
	server.engine.PATCH("reset-password", handler.ResetPasswordHandler)

	//player detail route
	server.engine.GET("/show-player", handler.ShowPlayerDetailsHandler)

	server.engine.GET("/trial", provider.PlayerAuthentication, handler.WelcomeHandler)

	//powerup routes
	server.engine.GET("/show-powerups", handler.ShowPowerUpsHandler)
	server.engine.POST("/use-powerup", provider.PlayerAuthentication, handler.UsePowerUpHandler)
	server.engine.POST("/buy-powerup", handler.BuyPowerupHandler)

	//leaderboard route
	server.engine.GET("/show-leaderboard", handler.ShowLeaderBoardHandler)

	//end game route
	server.engine.POST("/end-game", provider.PlayerAuthentication, handler.EndGameHandler)

	//swaggger route
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
