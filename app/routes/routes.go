package routes

import (
	"chess-move-api/app/controllers"

	"github.com/gin-gonic/gin"
)
func Routes(app *gin.Engine) {
	chessController := controllers.ChessController{}

	v1 := app.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(200, "Hi! Wellcome Chess Knight Move API")
		})
		v1.GET("/knight/:position",chessController.GetNextMoves) 
	}
}