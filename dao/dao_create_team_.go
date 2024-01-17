package dao

import (
    "context"
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"

)

func DB_CreateTeam (application *dto.Team) error {

	_, err := dbConfig.DATABASE.Collection("Teams").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}