package cache_manager

import "github.com/aws/aws-sdk-go/service/sqs"

func Initialize() {
	sqsSession := StartSQS()
	rdb := ConnectRedis()

	cm := CacheManager{
		SqsSession:     sqsSession,
		rdb:            rdb,
		publishedPosts: make(chan *sqs.Message),
	}

	go cm.Listen()

	go func() {
		for {
			message := <-cm.publishedPosts
			cm.HandlePublishedPost(message)
		}
	}()
}
