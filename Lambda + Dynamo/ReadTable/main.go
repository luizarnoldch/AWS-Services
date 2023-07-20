package main

import (
	"context"
	"fmt"
	"os"
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
    tableName := os.Getenv("TABLE_NAME")

	resp, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: event.ID,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("unable to get item, %v", err)
	}

	// Print the item's attribute values
	for k, v := range resp.Item {
		switch v := v.(type) {
		case *types.AttributeValueMemberS:
			fmt.Printf("%s: %s\n", k, v.Value)
		// Additional case clauses can be added here for other attribute value types
		default:
			fmt.Printf("Unexpected attribute value type for key: %s\n", k)
		}
	}

	return "Item fetched successfully", nil
}

func main() {
	lambda.Start(HandleRequest)
}
