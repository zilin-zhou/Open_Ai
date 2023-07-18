package db

import (
	"context"
	"fmt"
	"log"
	"open_ai_chat/db/tables"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// 创建一个表
func (cd ClientDynamoDB) CreateRegisterTable(TableName string) (*types.TableDescription, error) {
	cd.TableName = TableName
	var tableDesc *types.TableDescription
	table, err := cd.Client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("email"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("password"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("email"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("password"),
			KeyType:       types.KeyTypeRange,
		}},
		TableName: aws.String(cd.TableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Printf("Couldn't create table %v. Here's why: %v\n", cd.TableName, err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(cd.Client)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(cd.TableName)}, 5*time.Minute)
		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
		tableDesc = table.TableDescription
	}
	return tableDesc, err
}

// 列出所有表
func (cd ClientDynamoDB) ListTables() ([]string, error) {
	var tableNames []string
	tables, err := cd.Client.ListTables(
		context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("Couldn't list tables. Here's why: %v\n", err)
	} else {
		tableNames = tables.TableNames
	}
	return tableNames, err
}

// 向表中添加项目
func (cd ClientDynamoDB) AddRegisterInfo(register tables.Register, TableName string) error {
	cd.TableName = TableName
	item, err := attributevalue.MarshalMap(register)
	if err != nil {
		panic(err)
	}
	_, err = cd.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(cd.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

// 表中查询项目
func (cd ClientDynamoDB) Query(index, value, TableName string) ([]tables.Register, error) {
	var err error
	cd.TableName = TableName
	var response *dynamodb.QueryOutput
	var registerTbs []tables.Register
	keyEx := expression.Key(index).Equal(expression.Value(value))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	} else {
		response, err = cd.Client.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:                 aws.String(cd.TableName),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
		if err != nil {
			log.Printf("Couldn't query for movies released in %v. Here's why: %v\n", value, err)
		} else {
			err = attributevalue.UnmarshalListOfMaps(response.Items, &registerTbs)
			if err != nil {
				log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
			}
		}
	}
	return registerTbs, err
}

// 表中删除项目
func (cd ClientDynamoDB) DeleteUserInfo(register tables.Register, TableName string) error {
	cd.TableName = TableName
	_, err := cd.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(cd.TableName), Key: register.GetKey(),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", register.Email, err)
	}
	return err
}

// 更新表中项目
func (cd ClientDynamoDB) UpdateUserInfo(register tables.Register, username, TableName string) error {
	cd.TableName = TableName
	params, err := attributevalue.MarshalList([]interface{}{username, register.Email, register.Password})
	if err != nil {
		panic(err)
	}
	_, err = cd.Client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("UPDATE \"%v\" SET username=? WHERE email=? AND password=?",
				cd.TableName)),
		Parameters: params,
	})
	if err != nil {
		log.Printf("Couldn't update register %v. Here's why: %v\n", register.Email, err)
	}
	return err
}
