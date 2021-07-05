package Routes

import (
	"github.com/adiputra22/golang-gin-first/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("users", Controllers.GetUsers)
		v1.POST("user", Controllers.CreateUser)
		v1.GET("users/:id", Controllers.GetUserByID)
		v1.PUT("users/:id", Controllers.UpdateUser)
		v1.DELETE("users/:id", Controllers.DeleteUser)

	}

	return r
}
