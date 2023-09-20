package publisher

import (
	"encoding/json"
	"log"
	"os"
	"sqs-blog/models"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func (p *Publisher) HandleNewPost(m *sqs.Message) {
	var post models.NewPostMessage

	if err := json.Unmarshal([]byte(*m.Body), &post); err != nil {
		log.Printf("unmarshalling message: %s\n", err.Error())
		return
	}

	createdPost := CreatePost(post, p.db)

	_, err := p.SqsSession.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("NEW_POSTS_QUEUE_URL")),
		ReceiptHandle: m.ReceiptHandle,
	})

	if err != nil {
		log.Printf("deleting message: %s\n", err.Error())
	}

	postJson, err := json.Marshal(createdPost)

	if err != nil {
		log.Printf("marshalling post: %s\n", err.Error())
	}

	_, err = p.SqsSession.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(os.Getenv("PUBLISHED_POSTS_QUEUE_URL")),
		MessageBody: aws.String(string(postJson)),
	})

	if err != nil {
		log.Printf("sending message: %s\n", err.Error())
	}

	// log.Printf("message sent: %s\n", *m.MessageId)
}

func CreatePost(post models.NewPostMessage, db *gorm.DB) models.Post {
	postModel := models.Post{
		UID:     post.UID,
		Title:   post.Title,
		Content: post.Content,
		Slug:    slug.Make(post.Title + "-" + time.Now().Format(time.Stamp)),
	}
	if err := db.Create(&postModel).Error; err != nil {
		log.Printf("saving new post in database: %s\n", err.Error())
	}

	return postModel
}
