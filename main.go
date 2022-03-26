package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:ahmad@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connect eror")
	}

	fmt.Println("Database connected")

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRequest := book.BookRequest{
	// 	Title: "Tes Service",
	// 	Price: "20000",
	// }

	// bookService.Create(bookRequest)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/book", bookHandler.GetAll)
	v1.GET("/book/:id", bookHandler.Find)
	v1.POST("/book/store", bookHandler.Store)
	v1.PUT("/book/:id/update", bookHandler.Update)
	v1.DELETE("/book/:id/delete", bookHandler.Delete)
	router.Run()
}
