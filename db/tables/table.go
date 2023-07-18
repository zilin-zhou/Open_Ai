package tables

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

//用户注册表
type Register struct{
	Email string `dynamodbav:"email" json:"email" validate:"required,email"`   
	UserName string `dynamodbav:"username" json:"username"`
	Password string `dynamodbav:"password" json:"password" validate:"required,gte=8,lte=30"`
	// Age int `dynamodbav:"age" json:"age"`
	// Sex string `dynamodbav:"sex" json:"sex"`
}
// GetKey returns the composite primary key of the movie in a format that can be
// sent to DynamoDB.
func (reg Register) GetKey() map[string]types.AttributeValue {
	email, err := attributevalue.Marshal(reg.Email)
	if err != nil {
		panic(err)
	}
	password, err := attributevalue.Marshal(reg.Password)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"email": email, "password": password}
}
