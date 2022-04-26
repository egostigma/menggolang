package handler

import (
	"main/graph"
	"main/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func GraphQL() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
