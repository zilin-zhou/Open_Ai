package db

import (
	"context"
	"log"
	cf "open_ai_chat/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	region          = cf.Confs.GetString("aws_config.region")
	accessKeyId     = cf.Confs.GetString("aws_config.accessKeyId")
	secretAccessKey = cf.Confs.GetString("aws_config.secretAccessKey")
)

var DB *ClientDynamoDB

type ClientDynamoDB struct {
	Client    *dynamodb.Client
	TableName string
}

func init() {
	DB = NewDynamoClient()
}

func NewDynamoClient() *ClientDynamoDB {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		//config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId,secretAccessKey,"123456789")),
	)
	//sfmt.Println(region,accessKeyId,secretAccessKey)
	if err != nil {
		log.Fatalf("unable config:%s\n", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &ClientDynamoDB{
		Client: client,
	}
}
