package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)
var recipes []Recipe 

func init() {
	recipes = make([]Recipe, 0)
}

type Recipe struct {
	Id	string `json:"id"`
	Name string `json:"name"`
	Keywords []string `json:"keywords"`
	Ingredients []string `json:"Ingredients"`
	Instructions []string `json:"Instructions"`
	PublishedAt time.Time `json:"PublishedAt"`
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func NewRecipeHandler (c *gin.Context) {
var recipe Recipe

   if err := c.ShouldBindJSON(&recipe); err != nil {
	   c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	   })
	   return
   }

	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}
	func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run()

}