package main

import (
	"chess-move-api/app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.Routes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
		log.Printf("Servidor HTTP Rodando na porta - localhost:%s", port)
	}
	app.Run(":" + port)
}

