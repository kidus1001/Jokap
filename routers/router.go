package routers

import (
	"jokeapp/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetUp() *gin.Engine {
	r := gin.Default()

	api := r.Group("")
	{
		api.GET("/", controllers.Home)
		api.GET("/joke/:id", controllers.GetSpecificJoke)
		api.GET("/jokes", controllers.GetAllJokes)
	}
	return r
}
