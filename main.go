package main

import (
	"errors"
	"log"
	"sqs-blog/api"
	cache_manager "sqs-blog/cache-manager"
	"sqs-blog/publisher"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	publisher.Initialize()
	cache_manager.Initialize()

	api := api.NewApi()

	e := echo.New()
	e.POST("/post", func(c echo.Context) error {
		title := c.Request().PostFormValue("title")
		content := c.Request().PostFormValue("content")

		_, err := api.NewMessage(title, content)
		if err != nil {
			return c.String(500, err.Error())
		}
		return c.String(201, "Message Published!")

	})
	e.GET("/post/:slug", func(c echo.Context) error {
		post, err := api.GetPost(c.Param("slug"))
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return c.String(404, "not found")
			}
			return c.String(500, err.Error())
		}
		return c.JSON(200, post)
	})

	log.Println("Starting server at :3000")
	e.Logger.Fatal(e.Start(":3000"))
}
