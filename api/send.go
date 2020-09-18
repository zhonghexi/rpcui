package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Url   string      `form:"url"`
	Param interface{} `json: "param"`
}

func Send(ctx *gin.Context) {
	var req Data

	if ctx.ShouldBind(&req) == nil {
		log.Println(req.Param)
		fmt.Printf("类型：%T\n", req.Param)
	}

	// conn, err := grpc.Dial(":5001", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// c := pb.NewGreeterClient(conn)
	// r, err := c.Login(context.Background(), &rpc.LoginRequest{Username: u.Username, Password: u.Password})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"desc": "Success",
	})
}
