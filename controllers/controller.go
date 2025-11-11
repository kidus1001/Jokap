package controllers

import (
	"fmt"
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
		fmt.Println(response)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Error"})
	}
}

func GetAllJokes(c *gin.Context) {
	numStr := c.Query("num")
	if numStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Num parameter is required!"})
		return
	}

	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid num parameter"})
		return
	}

	jokes, flag := services.GetAllJokes(num)
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
