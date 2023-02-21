package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Recipe struct {
	Name string `json:"name"`
	Keywords []string `json:"keywords"`
	Ingredients []string `json:"Ingredients"`
	Instructions []string `json:"Instructions"`
	PublishedAt time.Time `json:"PublishedAt"`
}

func main() {
	router := gin.Default()
	router.Run()
}