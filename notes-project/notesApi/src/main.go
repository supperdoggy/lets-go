package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/newNote", newNote)
		api.POST("/updateNote", updateNote)
		api.POST("/share/:username", shareNote)
		api.POST("/getNotes", sendNotes)
	}
	if err := r.Run(":2020"); err != nil {
		fmt.Println(err.Error())
	}
}
