package publisher

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"gorm.io/gorm"
)

type Publisher struct {
	SqsSession *sqs.SQS
	db         *gorm.DB
	newPosts   chan *sqs.Message
}

func (p *Publisher) Listen() {
	for {
		output, err := p.SqsSession.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(os.Getenv("NEW_POSTS_QUEUE_URL")),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(20),
		})

		if err != nil {
			log.Printf("failed to fetch sqs message: %v", err)
		}

		for _, message := range output.Messages {
			p.newPosts <- message
		}
	}
}
