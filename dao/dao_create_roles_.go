package dao

import (
    "context"
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"

)

func DB_CreateRoles (application *dto.Roles) error {

	_, err := dbConfig.DATABASE.Collection("Roless").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}