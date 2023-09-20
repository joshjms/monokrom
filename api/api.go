package api

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sqs-blog/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/redis/go-redis/v9"
)

type API struct {
	rdb        *redis.Client
	SqsSession *sqs.SQS
}

func NewApi() *API {
	return &API{
		rdb:        ConnectRedis(),
		SqsSession: StartSQS(),
	}
}

func (a *API) NewMessage(title, content string) (string, error) {
	message := models.NewPostMessage{
		Title:   title,
		Content: content,
	}
	b, _ := json.Marshal(message)

	_, err := a.SqsSession.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    aws.String(os.Getenv("NEW_POSTS_QUEUE_URL")),
	})

	if err != nil {
		log.Print(err)
	}

	return "Message Published!", err
}

func (a *API) GetPost(slug string) (models.Post, error) {
	var p models.Post
	tr := a.rdb.Get(context.Background(), "post:"+slug)
	b, err := tr.Bytes()
	if err != nil {
		return models.Post{}, err
	}
	json.Unmarshal(b, &p)
	return p, nil
}
