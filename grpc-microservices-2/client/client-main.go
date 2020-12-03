package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	proto "grpc-microservices-2/proto"
)

func main() {
	conn, err:= grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	client:= proto.NewEmployeeServiceClient(conn)

	g:=gin.Default()

	g.GET("/example/:id/:name/:city", func(ctx *gin.Context){
		id := ctx.Param("id")
		name:= ctx.Param("name")
		city:= ctx.Param("city")

		data:= &proto.CreateEmployeeReq{ 
			Id: string(id),
			Name: string(name),
			City: string(city),
		}

		if resp, err := client.Employee(ctx, data); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(resp.Id, resp.Name, resp.City),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	
	if err:=g.Run(":8000"); err!=nil{
		log.Fatalf("Failed	to run server: %v", err)
	}
}
