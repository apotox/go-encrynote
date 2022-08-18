package queue

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue interface {
	Publish(id string) error
	Init() error
}

type QueueDeleteNote struct {
	Client *sqs.Client
}

func (qdn *QueueDeleteNote) Init() error {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalf("--------- unable to load SDK config, %v", err)
		panic(err)
	}

	qdn.Client = sqs.NewFromConfig(cfg)
	return nil
}

func (qdn *QueueDeleteNote) Publish(id string) error {

	out, err := qdn.Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: aws.String(id),
		QueueUrl:    aws.String(os.Getenv("QUEUE_URL")),
	})

	if err != nil {
		log.Printf("Error sending message: %s", err)
		return err
	}

	log.Printf("MessageId: %s", *out.MessageId)

	return nil

}

func GetQueueDeleteNote() *QueueDeleteNote {
	queue := new(QueueDeleteNote)
	queue.Init()
	return queue
}
