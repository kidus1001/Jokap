package controllers

import (
	"fmt"
	"jokeapp/models"
	"jokeapp/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to laugh :)"})
}

func GetSpecificJoke(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error converting string to int")
		return
	}

	response, flag := services.GetASpecificJoke(idInt)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Here are the jokes",
			"data":    response})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error"})
	}
}

func GetAllJokes(c *gin.Context) {
	var req models.JokeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
		return
	}
	jokes, flag := services.GetAllJokes(req.Num)
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    jokes,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong!",
		})
	}
}
