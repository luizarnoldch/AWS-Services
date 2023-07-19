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

	_, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: event.ID,
			},
		},
		ExpressionAttributeNames: map[string]string{
			"#N": "Name",
			"#E": "Email",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{
				Value: event.Name,
			},
			":e": &types.AttributeValueMemberS{
				Value: event.Email,
			},
		},
		UpdateExpression: aws.String("SET #N = :n, #E = :e"),
	})
	if err != nil {
		return "", fmt.Errorf("unable to update item, %v", err)
	}

	return "Item updated successfully", nil
}

func main() {
	lambda.Start(HandleRequest)
}
