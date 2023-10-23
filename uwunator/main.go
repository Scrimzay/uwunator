package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	rand.Seed(time.Now().UnixNano())

	router.GET("/", uwuGETHandler)
	router.POST("/", uwuPOSTHandler)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func uwuGETHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func uwuPOSTHandler(c *gin.Context) {
	userInput := c.PostForm("entry")
	userInput = strings.ToLower(userInput)
	words := strings.Fields(userInput)
	count := 0
	
	uwuMap := map[string]string{
		"r": "w",
		"l": "w",
		"h": "w",
	}

	facesOptions := []string{"UwU", "OwO", ">w<", "VwV", "^w^", ":3"}

	for i := 0; i < len(words); i++ {
		count++
		if count == 6 {
			randomNumber := rand.Intn(6)
			selectedFace := facesOptions[randomNumber]
			words = append(words[:i], append([]string{selectedFace}, words[i:]...)...)
			i++
			count = 0
		}
		word := words[i]
		for key, value := range uwuMap {
			word = strings.ReplaceAll(word, key, value)
		}
		words[i] = word
	}
	modifiedInput := strings.Join(words, " ")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Output": modifiedInput,
	})

}