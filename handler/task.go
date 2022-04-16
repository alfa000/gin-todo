package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	
}


func TaskHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"id" : "1",
		"title" : "Berlajar Go Dasar",
		"desc" : "gaskan",
		"assignee" : "Adnan",
	})
}