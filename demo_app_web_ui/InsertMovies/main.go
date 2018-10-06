package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	{
		ID:   1,
		Name: "Fight Club",
	},
	{
		ID:   2,
		Name: "Heat",
	},
	{
		ID:   3,
		Name: "Hulk",
	},
	{
		ID:   4,
		Name: "Thor",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie

	err := json.Unmarshal([]byte(req.Body), &movie)

	if err != nil {
		return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "Invalid payload",
			},
			nil
	}

	movies = append(movies, movie)

	response, err := json.Marshal(movies)

	if err != nil {
		return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err.Error(),
			},
			nil
	}

	return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(response),
		},
		nil
}

func main() {
	lambda.Start(insert)
}
