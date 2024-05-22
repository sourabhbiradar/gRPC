package main

import (
	"context"
	proto "ginAPI/protoc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.WithAPIClient

func main() {
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client = proto.NewWithAPIClient(conn)

	r := gin.Default()
	r.GET("/sent-message/:message", clientConnectionServer) // `:message` is dynamic varibale
	r.Run(":8000")
}

func clientConnectionServer(c *gin.Context) {
	msg := c.Param("message") // from `:message` to `msg`

	req := &proto.Creqt{Reqt: msg}

	client.ServerResp(context.TODO(), req)

	c.JSON(200, gin.H{
		"message": "message sent" + msg,
	})
}
