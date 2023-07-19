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
	ID string `json:"id"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	_, err = client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: event.ID,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("unable to delete item, %v", err)
	}

	return "Item deleted successfully", nil
}

func main() {
	lambda.Start(HandleRequest)
}
