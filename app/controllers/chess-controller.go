package controllers

import (
	datacontext "chess-move-api/data/context"
	"chess-move-api/repository"
	"context"
	"encoding/json"
	"strings"
	"time"

	"chess-move-api/domain/entities"

	"github.com/gin-gonic/gin"
)
var (
	collection = datacontext.GetDBContext("Chess-API-Log")
	ctx = context.Background()
)
type ChessController struct {
	knightrepository *repository.Knightrepository
	chessLogrepository *repository.ChessLogrepository
}

func (c *ChessController) GetNextMoves(ctx *gin.Context) {
	
	payload, _ := json.Marshal(ctx.Request.Header)
	position := string(payload)
	chesslog := entities.ChessLog{ Log: position, Title: "Request Header", Date: time.Now()}
	position = strings.ToUpper(ctx.Param("position"))
	
	c.chessLogrepository.AddLog(chesslog)

	nextMoves, err := c.knightrepository.GetNextsMoves(position)
	if err != nil {
		chesslog = entities.ChessLog{ Log: err.Error(), Title: "Error GetNextsMoves", Date: time.Now()}
		c.chessLogrepository.AddLog(chesslog)
		ctx.JSON(400, gin.H{
			"Status": false,
			"Code": 400,
			"Data": nil,
			"Message": err.Error(),
		})
		return
	}
	data := gin.H{
		"Status": true,
		"Code": 200,
		"Data": nextMoves.Moves,
		"Message": nil,
	}
	ctx.JSON(200, data)

	payload, _ = json.Marshal(nextMoves)

	chesslog = entities.ChessLog{ Log: string(payload) , Title: "Response Request - Position: " + position, Date: time.Now()}
	c.chessLogrepository.AddLog(chesslog)
}