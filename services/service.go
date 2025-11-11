rpackage services

import (
	"context"
	"encoding/json"
	"fmt"
	"jokeapp/models"
	"log"
	"strings"

	"google.golang.org/genai"
)

var jokes []models.Joke

func GetJokesAPI(num int) {
	apiKey := "Enter your ${API} key here"
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})

	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf(`Give me %d simple but funny jokes. The returned text should me a json format for my struct 
			type joke struct {
				Id int
				Content string
			}
			Int must be from 1 to 10
		`, num)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	response := strings.TrimSpace(result.Text())
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimPrefix(response, "\n")
	response = strings.TrimSuffix(response, "```")
	response = strings.TrimSpace(response)

	err = json.Unmarshal([]byte(response), &jokes)
	if err != nil {
		log.Fatal("Error while unmarshalling the json data", err)
	}
}

func GetAllJokes(num int) ([]models.Joke, bool) {
	GetJokesAPI(num)
	if len(jokes) == 0 {
		return []models.Joke{}, false
	} else {
		return jokes, true
	}
}

func GetASpecificJoke(id int) (models.Joke, bool) {
	if len(jokes) == 0 {
		fmt.Println("You havenot created jokes yet!")
		return models.Joke{}, false
	}
	return jokes[id], true
}
