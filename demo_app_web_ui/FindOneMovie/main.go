package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func findOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.PathParameters["id"])

	if err != nil {
		return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "ID must be a number",
			},
			nil
	}

	response, err := json.Marshal(movies[id-1])

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
	lambda.Start(findOne)
}
