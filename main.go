package main

import (
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gin-todo/handler"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	db.AutoMigrate(&task.Task{});

	router  := gin.Default();

	v1 := router.Group("/v1");

	v1.GET("/", rootHandler)
	v1.GET("/task", handler.TaskHandler)
	v1.GET("/task/:id", handler.TaskHandler) 

	router.Run();
}

func rootHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Adnan",
		"bio" : "Hallo guys!",
	})
}