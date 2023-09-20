package publisher

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Initialize() {
	db := StartDatabase()
	sqsSession := StartSQS()

	publisher := Publisher{
		SqsSession: sqsSession,
		db:         db,
		newPosts:   make(chan *sqs.Message),
	}

	go publisher.Listen()

	go func() {
		for {
			message := <-publisher.newPosts
			publisher.HandleNewPost(message)
		}
	}()
}
