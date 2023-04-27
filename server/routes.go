package server

import (
	_ "subway/docs"
	"subway/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//authentication route
	server.engine.POST("/register-player", handler.RegisterHandler)
	server.engine.POST("/login", handler.LoginHandler)
	server.engine.DELETE("/logout", handler.LogoutHandler)
	server.engine.PATCH("/update-pass", handler.UpdatePasswordHandler)
	server.engine.PATCH("/update-name", handler.UpdateNameHandler)

	//swaggger route
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
