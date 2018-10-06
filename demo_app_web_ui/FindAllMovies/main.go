package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var movies = []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}{
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

func findAll() (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(movies)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findAll)
}
