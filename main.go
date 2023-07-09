package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var db *gorm.DB

func main() {
	godotenv.Load(".env")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASS")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	apiPort := os.Getenv("API_PORT")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate the Book model
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
	}

	router := gin.Default()

	// Endpoint to insert a book
	router.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.Create(&book).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book inserted successfully"})
	})

	// Endpoint to list all books
	router.GET("/books", func(c *gin.Context) {
		var books []Book
		err := db.Find(&books).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"books": books})
	})

	// Start the API server
	err = router.Run(fmt.Sprintf(":%s", apiPort))
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
