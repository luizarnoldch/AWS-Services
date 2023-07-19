package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MyEvent struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("MyTable"),
		Item: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: event.ID,
			},
			"Name": &types.AttributeValueMemberS{
				Value: event.Name,
			},
			"Email": &types.AttributeValueMemberS{
				Value: event.Email,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("unable to put item, %v", err)
	}

	return "Item put successfully", nil
}

func main() {
	lambda.Start(HandleRequest)
}
