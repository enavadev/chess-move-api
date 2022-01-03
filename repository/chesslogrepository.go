package repository

import (
	"context"

	datacontext "chess-move-api/data/context"
	"chess-move-api/domain/entities"
)

var (
	collection = datacontext.GetDBContext("Chess-API-Log")
	ctx = context.Background()
)
type ChessLogrepository struct {

}
func (repo *ChessLogrepository) AddLog(cheslog entities.ChessLog) (bool, error) {

	var err error

	_, err = collection.InsertOne(ctx, cheslog)

	if err != nil {
		return false, err
	}

	return true, nil
}
