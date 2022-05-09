package main

import (
	"fmt"
	"log"
	"main/handler"
	"main/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8000"

type Person struct {
	Name string
	Age  int
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.Use(middleware.GinContextToContextMiddleware())
	r.POST("/query", handler.GraphQL())
	r.GET("/", handler.Playground())
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
