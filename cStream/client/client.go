package main

import (
	proto "cStream/protoc"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ClientStreamClient

func main() {
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	client = proto.NewClientStreamClient(conn)

	r := gin.Default()
	r.GET("/sent", clientConnectionServer)
	r.Run(":8008")
}

func clientConnectionServer(c *gin.Context) {
	reqts := []*proto.CReqt{ // stream of requests
		{Reqt: "Request 1"},
		{Reqt: "Request 2"},
		{Reqt: "Request 3"},
		{Reqt: "Request 4"},
	}

	stream, err := client.ServerReply(context.TODO()) // calling ServerReply()

	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}

	for _, rq := range reqts {
		err = stream.Send(rq) // sending requests if any errors handling them below
		if err != nil {
			fmt.Println("Request not fulfilled", err)
			return
		}
	}

	response, err := stream.CloseAndRecv() // server SendAndClose()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": response, // response from server dispalyed on web page
	})

}
