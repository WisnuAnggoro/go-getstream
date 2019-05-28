package main

import (
	"log"

	"github.com/gin-gonic/gin"
	stream "gopkg.in/GetStream/stream-go2.v3"

	"github.com/wisnuanggoro/go-getstream/config"
	"github.com/wisnuanggoro/go-getstream/getstream"
	"github.com/wisnuanggoro/go-getstream/handler"
)

func main() {
	// Get configuration
	cfg := config.Get()

	// Initialize Getstream Client
	getstreamClient, err := stream.NewClient(
		cfg.GoStreamAPIKey,
		cfg.GoStreamAPISecret,
		stream.WithAPIRegion(cfg.GoStreamAPIRegion),
	)
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	// Initialize services
	getstreamSvc := getstream.NewService(getstreamClient)

	// Initialize handlers
	getstreamHandler := handler.NewGetstreamHandler(getstreamSvc)

	// Initialize and run gin server
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Posting
		v1.POST("/post", getstreamHandler.AddPostByUserSerial)
		v1.GET("/post/:userSerial/summary", getstreamHandler.GetPostByUserSerial)
		v1.GET("/post/:userSerial/detail", getstreamHandler.GetPostDetailByUserSerial)
		v1.DELETE("/post", getstreamHandler.DeletePostByPostID)

		// Timeline
		v1.GET("/timeline/:userSerial/summary", getstreamHandler.GetTimelineByUserSerial)
		v1.GET("/timeline/:userSerial/detail", getstreamHandler.GetDetailTimelineByUserSerial)

		// Follower
		v1.POST("/user/follow", getstreamHandler.Follow)
		v1.POST("/user/unfollow", getstreamHandler.Unfollow)

		// Like Reaction
		v1.POST("/like", getstreamHandler.AddLikeToPostID)
		v1.GET("/like/:postID", getstreamHandler.RetrieveLikeDetailOnPostID)
		v1.GET("/like/:postID/:nextLikeID", getstreamHandler.RetrieveLikeDetailOnPostIDWithPagination)
		v1.DELETE("/like/:reactionID", getstreamHandler.RemoveLikeByReactionID)
	}
	router.Run()
}
