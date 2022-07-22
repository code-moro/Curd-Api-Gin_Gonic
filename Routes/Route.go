package Routes

import (
	"Rest-api/Controllers"
	"Rest-api/Database"

	"github.com/gin-gonic/gin"
)

func StartRouter(){
	Database.DatabaseConnect()
	router:=gin.Default()
    group:=router.Group("/book")
	{
	group.GET("/",Controllers.GetBooks)
	group.POST("/",Controllers.CreateBooks)
	group.GET("/:id",Controllers.GetBooksById)
	group.PUT("/:id",Controllers.UpdateBook)
	group.DELETE("/:id",Controllers.DeleteBook)
	}
	router.Run("localhost:8080")
}