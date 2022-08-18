package pkg

import (
	"github.com/apotox/go-encrynote/queue"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Datetime struct {
	primitive.DateTime
}

type Note struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	URL       string             `json:"url"`
	Message   string             `json:"message" validate:"min=1,max=140"`
	Password  string             `json:"password"`
	ExpireAt  Datetime           `json:"expireAt" bson:"expireAt"`
	CreatedAt Datetime           `json:"createdAt"`
}

type Services struct {
	QueueDeleteNote queue.Queue
}

type ResponseBodyType map[string]interface{}
