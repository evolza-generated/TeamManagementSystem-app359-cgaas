package dao

import (
    "context"
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"

)

func DB_CreateMember (application *dto.Member) error {

	_, err := dbConfig.DATABASE.Collection("Members").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}