package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khalidalhabibie/example_hello_world_grpc/grpc_gin/proto"
	"google.golang.org/grpc"
)

type Greeter struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.POST("/greeter", func(c *gin.Context) {
		var greeter1 Greeter

		if err := c.ShouldBindJSON(&greeter1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request Buddy"})
			return
		}

		name := greeter1.Name

		// Contact the server and print out its response.
		req := &pb.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}


