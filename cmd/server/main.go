package main

import (
	"log"
	"os"

	"auth-system/graph"
	"auth-system/internal/auth"
	"auth-system/internal/database"
	"auth-system/internal/middleware"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	database.ConnectDB()

	// Initialize Gin
	r := gin.Default()

	// Initialize auth service
	authService := auth.NewAuthService()

	// Create GraphQL handler
	resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	// Add auth middleware
	r.Use(middleware.AuthMiddleware(authService))

	// Routes
	r.POST("/query", gin.WrapH(srv))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(r.Run(":" + port))
}
