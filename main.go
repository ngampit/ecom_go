package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
   router := gin.Default()
   router.GET("/", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "Neversitup",
       })
   })
   router.Run()
	 loadenv()

	 // DB_USERNAME := os.Getenv("DB_USERNAME")
   // DB_NAME := os.Getenv("DB_NAME")
}