package main

import (
	"fmt"
	"log"
	"main/handler"
	"main/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Unable To Load .env File")
	}

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
