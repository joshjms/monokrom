package cache_manager

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sqs-blog/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/redis/go-redis/v9"
)

type CacheManager struct {
	SqsSession     *sqs.SQS
	rdb            *redis.Client
	publishedPosts chan *sqs.Message
}

func NewCacheManager(sqsSession *sqs.SQS, rdb *redis.Client) *CacheManager {
	return &CacheManager{
		SqsSession: sqsSession,
		rdb:        rdb,
	}
}

func (cm *CacheManager) Listen() {
	for {
		output, err := cm.SqsSession.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(os.Getenv("PUBLISHED_POSTS_QUEUE_URL")),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(20),
		})

		if err != nil {
			log.Printf("failed to fetch sqs message: %v", err)
		}

		for _, message := range output.Messages {
			cm.publishedPosts <- message
		}
	}
}

func (cm *CacheManager) HandlePublishedPost(m *sqs.Message) {
	var published models.PublishedPostMessage
	if err := json.Unmarshal([]byte(*m.Body), &published); err != nil {
		log.Printf("unmarshalling message: %s\n", err.Error())
		return
	}

	b, _ := json.Marshal(published.Post)
	cm.rdb.Set(context.Background(), "post:"+published.Slug, b, 0)

	_, err := cm.SqsSession.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("PUBLISHED_POSTS_QUEUE_URL")),
		ReceiptHandle: m.ReceiptHandle,
	})

	if err != nil {
		log.Printf("deleting message: %s\n", err.Error())
	}
}
