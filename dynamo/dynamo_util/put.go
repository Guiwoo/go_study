package dynamo_util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

func PutItem(ctx context.Context, c *dynamodb.Client, data interface{}) error {
	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}
	_, err = c.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}
