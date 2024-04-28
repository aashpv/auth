package router

import (
	"github.com/aashpv/auth/pkg/server/router/handlers"
	"github.com/gin-gonic/gin"
)

type Router interface {
	NewRouter() *gin.Engine
}

type router struct {
	hrs handlers.Handlers
}

func New(handlers handlers.Handlers) Router {
	return &router{hrs: handlers}
}

func (r *router) NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", r.hrs.SignUp)
	router.POST("/login", r.hrs.Login)

	//for middleware
	//router.Use(
	//	gin.Recovery(),
	//	gin.Logger(),
	//)
	//
	//auth := router.Group("/auth")
	//{
	//	auth.POST("/sign-up", r.hrs.SignUp)
	//	auth.POST("/sign-in", r.hrs.SignIn)
	//}
	return router
}
