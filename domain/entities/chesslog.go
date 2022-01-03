package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChessLog struct {
	ID 			  	primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Title		  	string				`bson:"Title,omitempty" json:"Title,omitempty"`
	Log			  	string				`bson:"Log,omitempty" json:"Log,omitempty"`
	Date		  	time.Time	    	`bson:"Date" json:"Date,omitempty"`
}